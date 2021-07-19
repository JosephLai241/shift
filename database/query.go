// Defining functions pertaining to querying the SQLite instance.

package database

import (
	"database/sql"
	"fmt"

	"github.com/JosephLai241/shift/utils"
)

// Struct used to store deserialized data from querying the SQLite instance.
type Deserialize struct {
	ShiftID         int
	Date            string
	Day             string
	ClockIn         string
	ClockInMessage  string
	ClockOut        string
	ClockOutMessage string
	ShiftDuration   string
	Month           string
}

// Run a SELECT SQL on the SQLite instance and return a slice of structs containing
// deserialized data.
func DeserializeRows(database *sql.DB, query string) []Deserialize {
	rows, err := database.Query(query)
	utils.CheckError(fmt.Sprintf("An error occurred when running query: %s", query), err)
	defer rows.Close()

	var deser Deserialize
	var dRows []Deserialize
	for rows.Next() {
		err := rows.Scan(
			&deser.ShiftID,
			&deser.Date,
			&deser.Day,
			&deser.ClockIn,
			&deser.ClockInMessage,
			&deser.ClockOut,
			&deser.ClockOutMessage,
			&deser.ShiftDuration,
			&deser.Month,
		)
		utils.CheckError("Failed to scan a row returned after querying the SQLite instance", err)

		dRows = append(dRows, deser)
	}

	return dRows
}
