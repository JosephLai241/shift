// Amend a shift in the timesheet.

package timesheet

import (
	"errors"
	"fmt"

	"github.com/JosephLai241/shift/utils"
	"github.com/JosephLai241/shift/views"
)

// Amend a clock-in or clock-out message within the timesheet.
func amendMessage(amendRow []string, newMessage string, target string) {
	var targetIndex int

	switch {
	case target == "in":
		targetIndex = 3
	case target == "out":
		targetIndex = 5
	}

	amendRow[targetIndex] = newMessage
}

// Amend the data in the selected shift and display the updated data.
func displayUpdate(args []string, rows [][]string, rowNums []int) ([]string, int) {
	intSelection := utils.CheckSelection(rowNums)
	amendRow := rows[intSelection]
	target := args[0]
	amendMessage(amendRow, args[1], target)

	fmt.Println("")
	utils.BoldWhite.Println("CHANGES")
	views.Display([][]string{amendRow})

	return amendRow, intSelection
}

// Amend the target shift.
func amendShift(amendRow []string, intSelection int, month string, rows [][]string, year string) {
	rows[intSelection] = amendRow

	overwriteTimesheet, err := utils.GetTimesheetByDFlags(month, true, year)
	if err != nil {
		utils.CheckError("Unable to open the timesheet to overwrite", err)
	}

	WriteToTimesheet(overwriteTimesheet, rows)
}

// Amend a shift in the timesheet.
func Amend(args []string, dayOrDate string, month string, year string) {
	timesheet, err := utils.GetTimesheetByDFlags(month, false, year)
	if err != nil {
		utils.CheckError(
			fmt.Sprintf("An error occurred when listing shifts recorded in %s %s", month, year),
			errors.New("no shifts were recorded"),
		)
	}

	rows := ReadTimesheet(timesheet)
	fmt.Println("")

	if rowNums, matches := findMatches(dayOrDate, rows); len(rowNums) == 0 {
		utils.CheckError("Error", fmt.Errorf("no shifts were found on %s", dayOrDate))
	} else {
		utils.BoldWhite.Println("MATCHES")
		views.DisplayOptions(matches)

		amendRow, intSelection := displayUpdate(args, rows, rowNums)

		switch confirmation := utils.ConfirmInput("revision"); confirmation {
		case "y":
			amendShift(amendRow, intSelection, month, rows, year)
			utils.BoldGreen.Printf("\nSuccessfully amended clock-%s message on %s.\n", args[0], dayOrDate)
		case "n":
			utils.BoldYellow.Printf("\nABORTING.\n")
		}
	}

	fmt.Println("")
}
