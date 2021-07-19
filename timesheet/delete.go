// Delete a shift in the timesheet.

package timesheet

import (
	"errors"
	"fmt"

	"github.com/JosephLai241/shift/utils"
	"github.com/JosephLai241/shift/views"
)

// Display the selected shift to be deleted from the timesheet.
func displayDeletion(rows [][]string, rowNums []int) int {
	intSelection := utils.CheckSelection(rowNums)
	rowForDeletion := rows[intSelection]

	fmt.Println("")
	utils.BoldRed.Println("PENDING DELETION")
	views.Display([][]string{rowForDeletion})

	return intSelection
}

// Remove the selected shift from the timesheet's rows.
func popShift(intSelection int, month string, rows [][]string, year string) {
	rows = append(rows[:intSelection], rows[intSelection+1:]...)

	overwriteTimesheet, err := utils.GetTimesheetByDFlags(month, true, year)
	if err != nil {
		utils.CheckError("Unable to open the timesheet to overwrite", err)
	}

	WriteToTimesheet(overwriteTimesheet, rows)
}

// Delete the selected shift from the timesheet.
func Delete(dayOrDate string, month string, year string) {
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

		intSelection := displayDeletion(rows, rowNums)
		switch confirmation := utils.ConfirmInput("deletion"); confirmation {
		case "y":
			popShift(intSelection, month, rows, year)
			utils.BoldGreen.Printf("\nSuccessfully deleted shift.\n")
		case "n":
			utils.BoldYellow.Printf("\nABORTING.\n")
		}
	}

	fmt.Println("")
}
