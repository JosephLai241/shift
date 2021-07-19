// List recorded shifts from querying the SQLite instance.

package database

import (
	"github.com/JosephLai241/shift/utils"
	"github.com/JosephLai241/shift/views"
)

// Pull and list records from the database.
func List(dayOrDate string, month string, subCommand string, year string) {
	database, err := OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	if dRows := queryMatches(database, dayOrDate, month, year); len(dRows) == 0 {
		utils.NoMatchesError(dayOrDate, month, year)
	} else {
		var shifts [][]string
		for _, row := range dRows {
			displayRow := []string{
				row.Date,
				row.Day,
				row.ClockIn,
				row.ClockInMessage,
				row.ClockOut,
				row.ClockOutMessage,
				row.ShiftDuration,
			}
			shifts = append(shifts, displayRow)
		}

		views.Display(shifts)
	}
}
