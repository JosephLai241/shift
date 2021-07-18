// Create, update, or delete shift data.

package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/JosephLai241/shift/modify"
	"github.com/JosephLai241/shift/utils"
)

// Initialize shift data.
type ShiftData struct {
	Type    string // "IN" or "OUT" (clock-in/out)
	Date    string // Shift (clock-in/out) date
	Day     string // Shift (clock-in/out) day
	Time    string // Shift (clock-in/out) time
	Message string // Complimentary message
}

// Calculate shift duration.
func (shiftData ShiftData) calculateDuration(startString string) string {
	startTime, _ := time.ParseInLocation("01-02-2006 15:04:05 Monday", startString, time.Now().Location())
	endTime, _ := time.ParseInLocation("01-02-2006 15:04:05 Monday", shiftData.Time, time.Now().Location())

	duration := endTime.Sub(startTime)

	return fmt.Sprint(duration)
}

// Timesheet functions.
// --------------------

// Write clock-in data to the timesheet. Add the column titles if the timesheet
// was just created.
func (shiftData ShiftData) recordInSheet(overwriteFile *os.File, rows [][]string) {
	if len(rows) == 0 {
		header := []string{
			"Date",
			"Day",
			"Clock-In",
			"Clock-In Message",
			"Clock-Out",
			"Clock-Out Message",
			"Shift Duration",
		}
		rows = append(rows, header)
	}

	newRow := []string{
		shiftData.Date,
		shiftData.Day,
		shiftData.Time,
		shiftData.Message,
		"TBD",
		"TBD",
		"TBD",
	}
	rows = append(rows, newRow)

	modify.WriteToTimesheet(overwriteFile, rows)
}

// Write clock-out data to the timesheet.
func (shiftData ShiftData) recordOutSheet(overwriteFile *os.File, rows [][]string) {
	for i, row := range rows {
		if i == len(rows)-1 {
			startString := fmt.Sprintf("%s %s %s", row[0], row[2], row[1])

			duration := shiftData.calculateDuration(startString)

			row[4] = shiftData.Time
			row[5] = shiftData.Message
			row[6] = duration
		}
	}

	modify.WriteToTimesheet(overwriteFile, rows)
}

// Write data stored in the ShiftData struct to the timesheet.
func (shiftData *ShiftData) RecordToTimesheet() {
	timesheetDirectory := modify.InitializeDirectories()
	timesheetPath := modify.GetTimesheetPath(timesheetDirectory)

	readFile := modify.InitializeTimesheet(timesheetPath)
	rows := modify.ReadTimesheet(readFile)

	overwriteFile, _ := os.OpenFile(timesheetPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if shiftData.Type == "IN" {
		shiftData.recordInSheet(overwriteFile, rows)
	} else {
		shiftData.recordOutSheet(overwriteFile, rows)
	}
}

// Find and return matches in the timesheet based on the day or date.
func FindMatches(dayOrDate string, rows [][]string) ([]int, [][]string) {
	var targetIndex int
	switch contains := strings.Contains(dayOrDate, "-"); contains {
	case true:
		targetIndex = 0
	case false:
		targetIndex = 1
	}

	var matches [][]string
	var rowNums []int
	for i, row := range rows {
		if row[targetIndex] == dayOrDate {
			matchRow := append([]string{strconv.Itoa(i)}, row...)
			matches = append(matches, matchRow)
			rowNums = append(rowNums, i)
		}
	}

	if len(matches) < 1 {
		return nil, nil
	}

	return rowNums, matches
}

// Amend a clock-in or clock-out message within the timesheet.
func AmendSheetMessage(amendRow []string, newMessage string, target string) {
	var targetIndex int

	switch {
	case target == "in":
		targetIndex = 3
	case target == "out":
		targetIndex = 5
	}

	amendRow[targetIndex] = newMessage
}

// SQLite functions.
// -----------------

// Format the message if there are quotes. Single quotes need to be doubled up
// (ie. `''`) when inserting into a SQL database.
func FormatMessage(message *string) {
	if strings.Contains(*message, "'") {
		*message = strings.ReplaceAll(*message, "'", "''")
	}
}

// Insert clock-in data into the SQLite instance.
func (shiftData ShiftData) recordInDB(database *sql.DB) {
	FormatMessage(&shiftData.Message)

	insertSQL := fmt.Sprintf(`
INSERT INTO M_%s (
	Date,
	Day,
	ClockIn,
	ClockInMessage,
	ClockOut,
	ClockOutMessage,
	ShiftDuration,
	Month
) VALUES (
	'%s',
	'%s',
	'%s',
	'%s',
	'%s',
	'%s',
	'%s',
	'%s'
);`,
		utils.CurrentMonth,
		shiftData.Date,
		shiftData.Day,
		shiftData.Time,
		shiftData.Message,
		"TBD",
		"TBD",
		"TBD",
		utils.CurrentMonth,
	)

	modify.ExecuteQuery(database, insertSQL)
}

// Update the last record in the SQLite instance with the clock-out data.
func (shiftData ShiftData) recordOutDB(database *sql.DB) {
	FormatMessage(&shiftData.Message)

	getLatestEntry := fmt.Sprintf(`
SELECT *
	FROM M_%s
	WHERE ShiftID = (SELECT MAX (ShiftID) FROM M_%s);
	`, utils.CurrentMonth, utils.CurrentMonth)

	dRows := modify.DeserializeRows(database, getLatestEntry)
	startString := dRows[0].Date + " " + dRows[0].ClockIn + " " + dRows[0].Day
	duration := shiftData.calculateDuration(startString)

	updateSQL := fmt.Sprintf(`
UPDATE M_%s
SET
	ClockOut = '%s',
	ClockOutMessage = '%s',
	ShiftDuration = '%s'
WHERE
	ShiftID = (SELECT MAX (ShiftID) FROM M_%s);
		`,
		utils.CurrentMonth,
		shiftData.Time,
		shiftData.Message,
		duration,
		utils.CurrentMonth,
	)

	modify.ExecuteQuery(database, updateSQL)
}

// Write data stored in the ShiftData struct to the SQLite instance.
func (shiftData *ShiftData) RecordToDB() {
	modify.StructureDB()

	database, err := modify.OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	if shiftData.Type == "IN" {
		shiftData.recordInDB(database)
	} else {
		shiftData.recordOutDB(database)
	}
}

// Query matches in the SQLite instance based on the day or date.
func QueryMatches(database *sql.DB, dayOrDate string, month string, year string) []modify.Deserialize {
	var target string
	switch contains := strings.Contains(dayOrDate, "-"); contains {
	case true:
		target = "Date"
	case false:
		target = "Day"
	}

	clause := fmt.Sprintf("HAVING %s = '%s'", target, dayOrDate)
	havingQuery := fmt.Sprintf(`
SELECT *
FROM M_%s
WHERE Month IN
	(SELECT Month FROM Y_%s WHERE Month = '%s')
GROUP BY ShiftID
%s;
	`,
		month,
		year,
		month,
		clause,
	)

	dRows := modify.DeserializeRows(database, havingQuery)
	if len(dRows) == 0 {
		utils.CheckError(
			fmt.Sprintf("Query: \n%s\n", havingQuery),
			errors.New("no records were found"),
		)
	}

	return dRows
}

// Amend the clock-in or clock-out message within the SQLite instance.
func AmendDBMessage(
	database *sql.DB,
	month string,
	shiftID string,
	newMessage string,
	target string,
	year string,
) {
	var setMessage string
	switch target {
	case "in":
		setMessage = fmt.Sprintf("ClockInMessage = '%s'", newMessage)
	case "out":
		setMessage = fmt.Sprintf("ClockOutMessage = '%s'", newMessage)
	}

	amendSQL := fmt.Sprintf(`
UPDATE M_%s
SET
	%s
WHERE Month IN
	(SELECT Month FROM Y_%s WHERE Month = '%s' AND ShiftID = '%s');
	`,
		month,
		setMessage,
		year,
		month,
		shiftID,
	)

	modify.ExecuteQuery(database, amendSQL)
}

// Delete a record in the SQLite instance.
func DeleteShiftDB(database *sql.DB, month string, shiftID string, year string) {
	deleteSQL := fmt.Sprintf(`
DELETE FROM M_%s
WHERE Month IN
	(SELECT Month FROM Y_%s WHERE Month = '%s' AND ShiftID = '%s');
	`,
		month,
		year,
		month,
		shiftID,
	)

	modify.ExecuteQuery(database, deleteSQL)
}
