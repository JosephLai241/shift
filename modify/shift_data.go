// Create, update, or delete shift data.

package modify

import (
	"fmt"
	"os"
	"time"
)

// Initialize shift data.
type ShiftData struct {
	Type    string // "IN" or "OUT" (clock-in/out)
	Date    string // Shift (clock-in/out) date
	Day     string // Shift (clock-in/out) day
	Time    string // Shift (clock-in/out) time
	Message string // Complimentary message
}

// Timesheet functions.

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

// SQLite functions.
