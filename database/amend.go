// Amend a shift in the SQLite instance.

package database

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"

	"github.com/JosephLai241/shift/utils"
	"github.com/JosephLai241/shift/views"
)

// Display the changes that will be made.
func displayUpdate(newMessage string, options [][]string, rowNum int, target string) {
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

// Amend the clock-in or clock-out message within the SQLite instance.
func amendMessage(
	database *sql.DB,
	month string,
	shiftID string,
	newMessage string,
	target string,
	year string,
) {
	var setMessage string
	switch target {
	case "in":
		setMessage = fmt.Sprintf("ClockInMessage = '%s'", newMessage)
	case "out":
		setMessage = fmt.Sprintf("ClockOutMessage = '%s'", newMessage)
	}

	amendSQL := fmt.Sprintf(`
UPDATE M_%s
SET
	%s
WHERE Month IN
	(SELECT Month FROM Y_%s WHERE Month = '%s' AND ShiftID = '%s');
	`,
		month,
		setMessage,
		year,
		month,
		shiftID,
	)

	ExecuteQuery(database, amendSQL)
}

// Amend a shift in the database.
func Amend(args []string, dayOrDate string, month string, year string) {
	database, err := OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	if dRows := queryMatches(database, dayOrDate, month, year); len(dRows) == 0 {
		utils.CheckError("Error", fmt.Errorf("no shifts were found on %s", dayOrDate))
	} else {
		options, rowNums := displayOptions(dRows)
		shiftID := utils.CheckSelection(rowNums)
		rowNum := sort.SearchInts(rowNums, shiftID)

		displayUpdate(args[1], options, rowNum, args[0])

		switch confirmation := utils.ConfirmInput("revision"); confirmation {
		case "y":
			FormatMessage(&args[1])
			amendMessage(database, month, strconv.Itoa(shiftID), args[1], args[0], year)
			utils.BoldGreen.Printf("\nSuccessfully amended clock-%s message on %s.\n", args[0], dayOrDate)
		case "n":
			utils.BoldYellow.Printf("\nABORTING.\n")
		}
	}

	fmt.Println("")
}
