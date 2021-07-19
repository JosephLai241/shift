// Initialize the SQLite instance.

package database

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
func insertYear(database *sql.DB) {
	insertYear := fmt.Sprintf(`
INSERT OR IGNORE INTO YEAR (Year)
VALUES (%s);
	`, utils.CurrentYear)

	ExecuteQuery(database, insertYear)
}

// Insert a new month into the current year's table.
func insertMonth(database *sql.DB) {
	insertMonth := fmt.Sprintf(`
INSERT OR IGNORE INTO Y_%s (Month, Year)
VALUES ('%s', %s);
	`, utils.CurrentYear, utils.CurrentMonth, utils.CurrentYear)

	ExecuteQuery(database, insertMonth)
}

// Create the .db binary and relational tables within it.
func StructureDB() {
	database, err := OpenDatabase()
	utils.CheckError("Could not open SQLite instance", err)
	defer database.Close()

	ExecuteQuery(database, "PRAGMA foreign_keys = ON;") // Enable foreign key constraint.
	ExecuteQuery(database, mainSQL)

	ExecuteQuery(database, yearSQL)
	insertYear(database)

	ExecuteQuery(database, monthSQL)
	insertMonth(database)
}
