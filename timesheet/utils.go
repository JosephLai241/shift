// Timesheet utilities.

package timesheet

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JosephLai241/shift/utils"
)

// Create the path to the timesheet.
func GetTimesheetPath(timesheetDirectory string) string {
	currentMonthYear := fmt.Sprintf("%s.csv", utils.CurrentMonth)
	timesheetPath := fmt.Sprintf("%s/%s", timesheetDirectory, currentMonthYear)

	return timesheetPath
}

// Find and return matches in the timesheet based on the day or date.
func findMatches(dayOrDate string, rows [][]string) ([]int, [][]string) {
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
