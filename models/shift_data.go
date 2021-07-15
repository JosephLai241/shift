// Create, update, or delete shift data.

package models

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/JosephLai241/shift/modify"
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
// --------------------

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

	modify.WriteToTimesheet(overwriteFile, rows)
}

// Write clock-out data to the timesheet.
func (shiftData ShiftData) recordOut(overwriteFile *os.File, rows [][]string) {
	for i, row := range rows {
		if i == len(rows)-1 {
			startString := fmt.Sprintf("%s %s %s", row[0], row[1], row[2])
			endString := fmt.Sprintf("%s %s", time.Now().Format("01-02-2006 Monday"), shiftData.Time)

			startTime, _ := time.Parse("01-02-2006 Monday 15:04:05", startString)
			endTime, _ := time.Parse("01-02-2006 Monday 15:04:05", endString)

			duration := endTime.Sub(startTime)

			row[4] = shiftData.Time
			row[5] = shiftData.Message
			row[6] = fmt.Sprint(duration)
		}
	}

	modify.WriteToTimesheet(overwriteFile, rows)
}

// Write data stored in the ShiftData struct to CSV.
func (shiftData *ShiftData) RecordShift() {
	timesheetDirectory := modify.InitializeDirectories()
	timesheetPath := modify.GetTimesheetPath(timesheetDirectory)

	readFile := modify.InitializeTimesheet(timesheetPath)
	rows := modify.ReadTimesheet(readFile)

	overwriteFile, _ := os.OpenFile(timesheetPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if shiftData.Type == "IN" {
		shiftData.recordIn(overwriteFile, rows)
	} else {
		shiftData.recordOut(overwriteFile, rows)
	}
}

// Find and return matches based on the day or date.
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

// Amend a clock-in or clock-out message.
func AmendMessage(amendRow []string, newMessage string, target string) {
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
