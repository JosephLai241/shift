// Database utilities.

package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/JosephLai241/shift/views"
)

// Format the message if there are quotes. Single quotes need to be doubled up
// (ie. `''`) when inserting into a SQL database.
func FormatMessage(message *string) {
	if strings.Contains(*message, "'") {
		*message = strings.ReplaceAll(*message, "'", "''")
	}
}

// Query matches in the SQLite instance based on the day or date.
func queryMatches(database *sql.DB, dayOrDate string, month string, year string) []Deserialize {
	var target string
	switch contains := strings.Contains(dayOrDate, "-"); contains {
	case true:
		target = "Date"
	case false:
		target = "Day"
	}

	clause := fmt.Sprintf("HAVING %s = '%s'", target, dayOrDate)
	havingQuery := fmt.Sprintf(`
SELECT *
FROM M_%s
WHERE Month IN
	(SELECT Month FROM Y_%s WHERE Month = '%s')
GROUP BY ShiftID
%s;
	`,
		month,
		year,
		month,
		clause,
	)

	dRows := DeserializeRows(database, havingQuery)

	return dRows
}

// Display options that were pulled from the query in a neat table.
func displayOptions(dRows []Deserialize) ([][]string, []int) {
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
