// Defining the `delete` command.

package cmd

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/modify"
	"github.com/JosephLai241/shift/utils"
	"github.com/JosephLai241/shift/views"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command.
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a recorded shift",
	Long: `
   _     _     _       
 _| |___| |___| |_ ___ 
| . | -_| | -_|  _| -_|
|___|___|_|___|_| |___|

Use this command to delete a recorded shift. This command 
is fairly versatile - you can search for records based on 
a day of the week or date, month, and/or year.

Using delete without additional commands or flags will display a
table containing shifts recorded for the current day.

There are three optional flags you can use: the '-d', '-m', 
and '-y' flags. These flags denote the target day of the week or date, 
month, and year respectively. The default value for all of these 
flags is the current day of the week/date, month, and year. 
Combine these flags to to do a deep search for a particular 
shift or shifts.

You can search for shifts on a different day or date by using the '-d'
flag, which accepts a day of the week (ie. Monday) or a date 
(ie. 07-14-2021). The accepted date formats are:

- MM-DD-YYYY
- MM/DD/YYYY

You can search for shifts in a different month by using the
'-m' flag, which accepts a month (ie. July). If this is the only
provided flag, a search will be done for the current day within
the provided month.

Finally, you can search for shifts in a different year by using
the '-y' flag, which accepts a year (ie. 2021). Like the '-m'
flag, a search will be done for the current day and month within
the provided year if this is the only provided flag.

You can combine the '-d', '-m', and/or '-y' flags to do a deep
search for a particular shift or shifts.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.DeleteArt)

		dayOrDate, month, year := formatFlags(cmd)
		month = strings.Title(month)

		modify.CRUD(
			func() { deleteShiftTimesheet(dayOrDate, month, year) },
			func() { deleteShiftDatabase(dayOrDate, month, year) },
		)
	},
}

// Add the `delete` command and its sub-flags to the base command.
func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP(
		"dayordate", "d",
		utils.CurrentDate,
		"Search records on a day of the week or date",
	)
	deleteCmd.Flags().StringP(
		"month", "m",
		utils.CurrentMonth,
		"Search records in a month",
	)
	deleteCmd.Flags().StringP(
		"year", "y",
		utils.CurrentYear,
		"Search records in a year",
	)
}

// Timesheet functions.
// --------------------

// Display the selected shift to be deleted from the timesheet.
func displayDeletion(rows [][]string, rowNums []int) int {
	intSelection := checkSelection(rowNums)
	rowForDeletion := rows[intSelection]

	fmt.Println("")
	utils.BoldRed.Println("PENDING DELETION")
	views.Display([][]string{rowForDeletion})

	return intSelection
}

// Remove the selected shift from the timesheet's rows.
func popShift(intSelection int, month string, rows [][]string, year string) {
	rows = append(rows[:intSelection], rows[intSelection+1:]...)

	overwriteTimesheet, err := getTimesheetByDFlags(month, true, year)
	if err != nil {
		utils.CheckError("Unable to open the timesheet to overwrite", err)
	}

	modify.WriteToTimesheet(overwriteTimesheet, rows)
}

// Delete the selected shift from the timesheet.
func deleteShiftTimesheet(dayOrDate string, month string, year string) {
	timesheet, err := getTimesheetByDFlags(month, false, year)
	if err != nil {
		utils.CheckError(
			// fmt.Sprintf("An error occurred when listing shifts recorded in %s %s", strings.Title(month), year),
			fmt.Sprintf("An error occurred when listing shifts recorded in %s %s", month, year),
			errors.New("no shifts were recorded"),
		)
	}

	rows := modify.ReadTimesheet(timesheet)
	fmt.Println("")

	switch rowNums, matches := models.FindMatches(dayOrDate, rows); len(rowNums) {
	case 0:
		utils.CheckError("Error", fmt.Errorf("no shifts were found on %s", dayOrDate))
	default:
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

// SQLite functions.
// -----------------

// Delete the selected shift from the database.
func deleteShiftDatabase(dayOrDate string, month string, year string) {
	database, err := modify.OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	dRows := models.QueryMatches(database, dayOrDate, strings.Title(month), year)

	options, rowNums := displayDBOptions(dRows)
	for i, row := range options {
		options[i] = row[1:]
	}

	shiftID := checkSelection(rowNums)
	rowNum := sort.SearchInts(rowNums, shiftID)

	fmt.Println("")
	utils.BoldRed.Println("PENDING DELETION")
	views.Display([][]string{options[rowNum]})

	switch confirmation := utils.ConfirmInput("deletion"); confirmation {
	case "y":
		models.DeleteShiftDB(database, month, strconv.Itoa(shiftID), year)
		utils.BoldGreen.Printf("\nSuccessfully deleted shift.\n")
	case "n":
		utils.BoldYellow.Printf("\nABORTING.\n")
	}

	fmt.Println("")
}
