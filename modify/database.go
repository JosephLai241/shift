// Write shift data to a SQLite instance.

package modify

import (
	"database/sql"
	"fmt"

	"github.com/JosephLai241/shift/utils"
	_ "github.com/mattn/go-sqlite3"
)

// SQL for creating the main "YEAR" table.
var mainSQL string = `
CREATE TABLE IF NOT EXISTS YEAR (
	MainID INTEGER PRIMARY KEY AUTOINCREMENT, 
	Year TEXT,
	UNIQUE (Year)
);`

// SQL for creating the current year's table.
var yearSQL string = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS Y_%s (
	YearID INTEGER PRIMARY KEY AUTOINCREMENT,
	Month TEXT,
	Year TEXT,
	FOREIGN KEY (Year) REFERENCES YEAR (Year),
	UNIQUE (Month)
);`,
	utils.CurrentYear,
)

// SQL for creating the current month's table.
var monthSQL string = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS M_%s (
	ShiftID INTEGER PRIMARY KEY AUTOINCREMENT,
	Date TEXT,
	Day TEXT,
	ClockIn TEXT,
	ClockInMessage TEXT,
	ClockOut TEXT NOT NULL,
	ClockOutMessage TEXT NOT NULL,
	ShiftDuration TEXT NOT NULL,
	Month TEXT NOT NULL,
	FOREIGN KEY (Month) REFERENCES Y_%s (Month)
);`,
	utils.CurrentMonth,
	utils.CurrentYear,
)

// Open the `shifts.db` SQLite instance.
func OpenDatabase() (*sql.DB, error) {
	cwd := utils.GetCWD()
	databasePath := fmt.Sprintf("%s/shifts.db", cwd)
	database, err := sql.Open("sqlite3", databasePath)

	return database, err
}

// Insert a new year into the `YEAR` table.
func InsertYear(database *sql.DB) {
	// SQL for inserting a new year.
	insertYear := fmt.Sprintf(`
		INSERT OR IGNORE INTO YEAR (Year)
		VALUES (%s);
	`, utils.CurrentYear)

	ExecuteQuery(database, insertYear)
}

// Insert a new month into the current year's table.
func InsertMonth(database *sql.DB) {
	// SQL for inserting a new month.
	insertMonth := fmt.Sprintf(`
		INSERT OR IGNORE INTO Y_%s (Month, Year)
		VALUES ('%s', %s);
	`, utils.CurrentYear, utils.CurrentMonth, utils.CurrentYear)

	ExecuteQuery(database, insertMonth)
}

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

// Query the SQLite instance and get all shifts from a `M_MONTH` table.
func QueryTable(month string) []Deserialize {
	database, err := OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	rows, _ := database.Query(fmt.Sprintf("SELECT * FROM M_%s", month))
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

// Execute a query within the SQLite instance.
func ExecuteQuery(database *sql.DB, query string) {
	statement, err := database.Prepare(query)
	utils.CheckError(fmt.Sprintf("An error occurred when running the SQL command: \n%s\n\n", query), err)
	statement.Exec()
}

func StructureDB() {
	database, err := OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	ExecuteQuery(database, "PRAGMA foreign_keys = ON;") // Enable foreign key constraint.
	ExecuteQuery(database, mainSQL)

	ExecuteQuery(database, yearSQL)
	InsertYear(database)

	ExecuteQuery(database, monthSQL)
	InsertMonth(database)
}
