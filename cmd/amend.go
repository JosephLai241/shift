// Defining the `amend` command.

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

// amendCmd represents the amend command.
var amendCmd = &cobra.Command{
	Use:   "amend (in|out)",
	Short: "Amend a shift's clock-in or clock-out message",
	Long: `
                     _ 
 ___ _____ ___ ___ _| |
| .'|     | -_|   | . |
|__,|_|_|_|___|_|_|___|

Use this command to amend a recorded shift's clock-in or clock-out
message. This command is fairly versatile - you can search for records 
based on a day of the week or date, month, and/or year.

Using amend without additional commands or flags will display a
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
		fmt.Println(utils.AmendArt)

		checkArgs(args)
		args[0] = strings.ToLower(args[0])
		dayOrDate, month, year := formatFlags(cmd)

		modify.CRUD(
			func() { amendTimesheet(args, dayOrDate, month, year) },
			func() { amendDatabase(args, dayOrDate, month, year) },
		)
	},
}

// Add the `amend` command and its sub-flags to the base command.
func init() {
	rootCmd.AddCommand(amendCmd)

	amendCmd.Flags().StringP(
		"dayordate", "d",
		utils.CurrentDate,
		"Search records on a day of the week or date",
	)
	amendCmd.Flags().StringP(
		"month", "m",
		utils.CurrentMonth,
		"Search records in a month",
	)
	amendCmd.Flags().StringP(
		"year", "y",
		utils.CurrentYear,
		"Search records in a year",
	)
}

// Check all input for the `amend` command.
func checkArgs(args []string) {
	if len(args) < 1 {
		utils.CheckError("Command error", errors.New("`amend` requires in or out"))
	} else if len(args) < 2 {
		utils.CheckError("Command error", errors.New("`amend` requires a new message"))
	} else {
		utils.BoldBlue.Printf("New message: %s\n", args[len(args)-1])
	}
}

// Check if the selection is valid.
func checkSelection(rowNums []int) int {
	validOptions := make(map[int]struct{})
	for i := range rowNums {
		validOptions[rowNums[i]] = struct{}{}
	}

	var input string
	for {
		fmt.Printf("\nSelect a shift to modify %+v: ", rowNums)
		fmt.Scanln(&input)

		intSelection, _ := strconv.Atoi(input)
		if _, ok := validOptions[intSelection]; !ok {
			utils.BoldRed.Print("\nInvalid option. Try again.\n")
		} else {
			return intSelection
		}
	}
}

// Timesheet functions.
// --------------------

// Amend the data in the selected shift and display the updated data.
func displaySheetUpdate(args []string, rows [][]string, rowNums []int) ([]string, int) {
	intSelection := checkSelection(rowNums)
	amendRow := rows[intSelection]
	target := args[0]
	models.AmendSheetMessage(amendRow, args[1], target)

	fmt.Println("")
	utils.BoldWhite.Println("CHANGES")
	views.Display([][]string{amendRow})

	return amendRow, intSelection
}

// Amend the target shift.
func amendShift(amendRow []string, intSelection int, month string, rows [][]string, year string) {
	rows[intSelection] = amendRow

	overwriteTimesheet, err := getTimesheetByDFlags(month, true, year)
	if err != nil {
		utils.CheckError("Unable to open the timesheet to overwrite", err)
	}

	modify.WriteToTimesheet(overwriteTimesheet, rows)
}

// Amend a shift in the timesheet.
func amendTimesheet(args []string, dayOrDate string, month string, year string) {
	timesheet, err := getTimesheetByDFlags(month, false, year)
	if err != nil {
		utils.CheckError(
			fmt.Sprintf("An error occurred when listing shifts recorded in %s %s", strings.Title(month), year),
			errors.New("no shifts were recorded"),
		)
	}

	rows := modify.ReadTimesheet(timesheet)
	fmt.Println("")

	if rowNums, matches := models.FindMatches(dayOrDate, rows); len(rowNums) == 0 {
		utils.CheckError("Error", fmt.Errorf("no shifts were found on %s", dayOrDate))
	} else {
		utils.BoldWhite.Println("MATCHES")
		views.DisplayOptions(matches)

		amendRow, intSelection := displaySheetUpdate(args, rows, rowNums)

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

// SQLite functions.
// -----------------

// Display options that were pulled from the query in a neat table.
func displayDBOptions(dRows []modify.Deserialize) ([][]string, []int) {
	var rowNums []int
	var options [][]string
	for _, row := range dRows {
		rowNums = append(rowNums, row.ShiftID)

		shiftID := strconv.Itoa(row.ShiftID)
		displayRow := []string{
			shiftID,
			row.Date,
			row.Day,
			row.ClockIn,
			row.ClockInMessage,
			row.ClockOut,
			row.ClockOutMessage,
			row.ShiftDuration,
		}
		options = append(options, displayRow)
	}

	views.DisplayOptions(options)

	return options, rowNums
}

// Display the changes that will be made.
func displayDBUpdate(newMessage string, options [][]string, rowNum int, target string) {
	targetRow := options[rowNum]
	targetRow = targetRow[1:]

	var targetIndex int
	switch target {
	case "in":
		targetIndex = 3
	case "out":
		targetIndex = 5
	}

	targetRow[targetIndex] = newMessage

	utils.BoldWhite.Println("CHANGES")
	views.Display([][]string{targetRow})
}

// Amend a shift in the database.
func amendDatabase(args []string, dayOrDate string, month string, year string) {
	database, err := modify.OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	if dRows := models.QueryMatches(database, dayOrDate, month, year); len(dRows) == 0 {
		utils.CheckError("Error", fmt.Errorf("no shifts were found on %s", dayOrDate))
	} else {
		options, rowNums := displayDBOptions(dRows)
		shiftID := checkSelection(rowNums)
		rowNum := sort.SearchInts(rowNums, shiftID)

		displayDBUpdate(args[1], options, rowNum, args[0])

		switch confirmation := utils.ConfirmInput("revision"); confirmation {
		case "y":
			models.FormatMessage(&args[1])
			models.AmendDBMessage(database, month, strconv.Itoa(shiftID), args[1], args[0], year)
			utils.BoldGreen.Printf("\nSuccessfully amended clock-%s message on %s.\n", args[0], dayOrDate)
		case "n":
			utils.BoldYellow.Printf("\nABORTING.\n")
		}
	}

	fmt.Println("")
}
