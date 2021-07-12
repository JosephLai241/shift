// Write shift data to a SQLite instance.

package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/JosephLai241/shift/utils"
	_ "github.com/mattn/go-sqlite3"
)

var currentYear string = time.Now().Format("2006")
var currentmonth string = time.Now().Format("January")

// SQL for creating the main "YEAR" table.
var mainSQL string = `
CREATE TABLE YEAR (
	MainID INTEGER PRIMARY KEY, 
	Year TEXT
);`

// SQL for creating the current year's table.
var yearSQL string = fmt.Sprintf(`
CREATE TABLE Y_%s (
	YearID INTEGER PRIMARY KEY,
	Month TEXT,
	MainID INTEGER,
	FOREIGN KEY (MainID) REFERENCES YEAR (MainID)
);`,
	currentYear,
)

// SQL for creating the current month's table.
var monthSQL string = fmt.Sprintf(`
CREATE TABLE M_%s (
	MonthID INTEGER PRIMARY KEY,
	Date TEXT,
	Day TEXT,
	ClockIn TEXT,
	ClockInMessage TEXT,
	ClockOut TEXT,
	ClockOutMessage TEXT,
	ShiftDuration TEXT,
	YearID INTEGER,
	FOREIGN KEY(YearID) REFERENCES Y_%s(YearID)
);`,
	currentmonth,
	currentYear,
)

// Spawn a SQLite instance where shift data is stored.
func spawnDatabase() *sql.DB {
	cwd := utils.GetCWD()
	databasePath := fmt.Sprintf("%s/shifts.db", cwd)
	database, err := sql.Open("sqlite3", databasePath)
	utils.CheckError("Could not spawn SQLite instance", err)

	return database
}

// Create a table within the SQLite instance.
func createTable(database *sql.DB, query string) {
	statement, err := database.Prepare(query)
	utils.CheckError(fmt.Sprintf("An error occurred when trying to create a table. SQL is: \n%s\n\n", query), err)
	statement.Exec()
}

func WriteToDB() {
	database := spawnDatabase()
	defer database.Close()

	createTable(database, mainSQL)
	createTable(database, yearSQL)
	createTable(database, monthSQL)

}
