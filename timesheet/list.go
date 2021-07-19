// List recorded shifts pulled from the timesheet.

package timesheet

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JosephLai241/shift/utils"
	"github.com/JosephLai241/shift/views"
)

// Find matches based on search day or date and format the records.
func getAndFormatMatches(dayOrDate string, rows [][]string) [][]string {
	_, matches := findMatches(dayOrDate, rows)
	if len(matches) == 0 {
		utils.CheckError(
			"Error",
			errors.New("no shifts were found based on your search parameters"),
		)
	}

	for i, row := range matches {
		matches[i] = row[1:]
	}

	return matches
}

// List matches.
func listMatches(dayOrDate string, matches [][]string, month string, year string) {
	if strings.Contains(dayOrDate, "-") {
		utils.BoldBlue.Printf("Displaying shifts recorded on %s.\n\n", dayOrDate)
	} else {
		utils.BoldBlue.Printf("Displaying shifts recorded on a %s within %s %s.\n\n", dayOrDate, month, year)
	}

	views.Display(matches)
}

// Pull and list records from timesheets.
func List(dayOrDate string, month string, subCommand string, year string) {
	timesheet, err := utils.GetTimesheetByDFlags(month, false, year)
	if err != nil {
		utils.CheckError(
			fmt.Sprintf("An error occurred when listing shifts recorded in %s %s", month, year),
			errors.New("no shifts were recorded"),
		)
	}

	if rows := ReadTimesheet(timesheet); len(rows) == 0 {
		utils.NoMatchesError(dayOrDate, month, year)
	} else {
		rows = rows[1:]

		if subCommand == "all" {
			utils.BoldBlue.Printf("Displaying all shifts recorded in %s %s.\n\n", month, year)
			views.Display(rows)
		} else {
			matches := getAndFormatMatches(dayOrDate, rows)
			listMatches(dayOrDate, matches, month, year)
		}
	}
}
