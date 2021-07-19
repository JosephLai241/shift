// Delete a shift in the SQLite instance.

package database

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"

	"github.com/JosephLai241/shift/utils"
	"github.com/JosephLai241/shift/views"
)

// Delete a single record in the SQLite instance.
func deleteSingle(database *sql.DB, month string, shiftID string, year string) {
	deleteSQL := fmt.Sprintf(`
DELETE FROM M_%s
WHERE Month IN
	(SELECT Month FROM Y_%s WHERE Month = '%s' AND ShiftID = '%s');
	`,
		month,
		year,
		month,
		shiftID,
	)

	ExecuteQuery(database, deleteSQL)
}

// Delete the selected shift from the database.
func Delete(dayOrDate string, month string, year string) {
	database, err := OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	if dRows := queryMatches(database, dayOrDate, month, year); len(dRows) == 0 {
		utils.CheckError("Error", fmt.Errorf("no shifts were found on %s", dayOrDate))
	} else {
		options, rowNums := displayOptions(dRows)
		for i, row := range options {
			options[i] = row[1:]
		}

		shiftID := utils.CheckSelection(rowNums)
		rowNum := sort.SearchInts(rowNums, shiftID)

		fmt.Println("")
		utils.BoldRed.Println("PENDING DELETION")
		views.Display([][]string{options[rowNum]})

		switch confirmation := utils.ConfirmInput("deletion"); confirmation {
		case "y":
			deleteSingle(database, month, strconv.Itoa(shiftID), year)
			utils.BoldGreen.Printf("\nSuccessfully deleted shift.\n")
		case "n":
			utils.BoldYellow.Printf("\nABORTING.\n")
		}
	}

	fmt.Println("")
}
