// Write shift data to the `CURRENT_MONTH.csv` timesheet.

package modify

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/JosephLai241/shift/utils"
)

// Initialize the directories in which the timesheet will be stored.
// Returns a string denoting the path to the timesheet directory.
func initializeDirectories() string {
	cwd := GetCWD()
	timesheetDirectory := fmt.Sprintf("%s/shift_timesheets/%s", cwd, time.Now().Format("2006"))
	os.MkdirAll(timesheetDirectory, os.ModePerm)

	return timesheetDirectory
}

// Create the path to the timesheet.
func getTimesheetPath(timesheetDirectory string) string {
	currentMonthYear := fmt.Sprintf("%s.csv", time.Now().Format("January"))
	timesheetPath := fmt.Sprintf("%s/%s", timesheetDirectory, currentMonthYear)

	return timesheetPath
}

// If the file does not already exist, create a new timesheet with column headers.
// Returns *os.File for writing.
func initializeTimesheet(timesheetPath string) *os.File {
	var file *os.File
	if _, err := os.Stat(timesheetPath); err == nil {
		file, _ = os.OpenFile(timesheetPath, os.O_RDWR, 0755)
		utils.CheckError("Could not open timesheet", err)
	} else if os.IsNotExist(err) {
		file, err = os.Create(timesheetPath)
		utils.CheckError("Could not create timesheet", err)
	}

	return file
}

// Write data to to the timesheet.
func writeToTimesheet(file *os.File, rows [][]string) {
	writer := csv.NewWriter(file)
	err := writer.WriteAll(rows)
	utils.CheckError("Could not flush writer data after writing data to timesheet", err)

	closeErr := file.Close()
	utils.CheckError("Could not write shift data to timesheet", closeErr)
}

// Read timesheet and extract the data into a nested list of strings.
func readTimesheet(file *os.File) [][]string {
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	utils.CheckError("Could not read timesheet", err)

	file.Close()
	return rows
}

// Initialize shift data.
type ShiftData struct {
	Type    string // "IN" or "OUT" (clock-in/out)
	Date    string // Shift (clock-in/out) date
	Day     string // Shift (clock-in/out) day
	Time    string // Shift (clock-in/out) time
	Message string // Complimentary message
}

// Write clock-in data to the timesheet. Add the column titles if the timesheet
// was just created.
func (shiftData ShiftData) recordIn(overwriteFile *os.File, rows [][]string) {
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
		"",
		"",
		"",
	}
	rows = append(rows, newRow)
	writeToTimesheet(overwriteFile, rows)
}

// Write clock-out data to the timesheet.
func (shiftData ShiftData) recordOut(overwriteFile *os.File, rows [][]string) {
	for i, row := range rows {
		if i == len(rows)-1 {
			startTime, _ := time.Parse("15:04:05", row[2])
			endTime, _ := time.Parse("15:04:05", shiftData.Time)
			duration := endTime.Sub(startTime)

			row[4] = shiftData.Time
			row[5] = shiftData.Message
			row[6] = fmt.Sprint(duration)
		}
	}

	writeToTimesheet(overwriteFile, rows)
}

// Write data stored in the ShiftData struct to CSV.
func (shiftData *ShiftData) RecordShift() {
	timesheetDirectory := initializeDirectories()
	timesheetPath := getTimesheetPath(timesheetDirectory)

	readFile := initializeTimesheet(timesheetPath)
	rows := readTimesheet(readFile)

	overwriteFile, _ := os.OpenFile(timesheetPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if shiftData.Type == "IN" {
		shiftData.recordIn(overwriteFile, rows)
	} else {
		shiftData.recordOut(overwriteFile, rows)
	}
}
