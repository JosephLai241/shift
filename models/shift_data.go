// Create, update, or delete shift data.

package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/JosephLai241/shift/database"
	"github.com/JosephLai241/shift/timesheet"
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

	timesheet.WriteToTimesheet(overwriteFile, rows)
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

	timesheet.WriteToTimesheet(overwriteFile, rows)
}

// Write data stored in the ShiftData struct to the timesheet.
func (shiftData *ShiftData) RecordToTimesheet() {
	timesheetDirectory := timesheet.InitializeDirectories()
	timesheetPath := timesheet.GetTimesheetPath(timesheetDirectory)

	readFile := timesheet.InitializeTimesheet(timesheetPath)
	rows := timesheet.ReadTimesheet(readFile)

	overwriteFile, _ := os.OpenFile(timesheetPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if shiftData.Type == "IN" {
		shiftData.recordInSheet(overwriteFile, rows)
	} else {
		shiftData.recordOutSheet(overwriteFile, rows)
	}
}

// SQLite functions.
// -----------------

// Insert clock-in data into the SQLite instance.
func (shiftData ShiftData) recordInDB(db *sql.DB) {
	database.FormatMessage(&shiftData.Message)

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

	database.ExecuteQuery(db, insertSQL)
}

// Update the last record in the SQLite instance with the clock-out data.
func (shiftData ShiftData) recordOutDB(db *sql.DB) {
	database.FormatMessage(&shiftData.Message)

	getLatestEntry := fmt.Sprintf(`
SELECT *
	FROM M_%s
	WHERE ShiftID = (SELECT MAX (ShiftID) FROM M_%s);
	`, utils.CurrentMonth, utils.CurrentMonth)

	dRows := database.DeserializeRows(db, getLatestEntry)
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

	database.ExecuteQuery(db, updateSQL)
}

// Write data stored in the ShiftData struct to the SQLite instance.
func (shiftData *ShiftData) RecordToDB() {
	database.StructureDB()

	database, err := database.OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	if shiftData.Type == "IN" {
		shiftData.recordInDB(database)
	} else {
		shiftData.recordOutDB(database)
	}
}
