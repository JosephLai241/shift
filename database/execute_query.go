// Execute a query on the SQLite instance.

package database

import (
	"database/sql"
	"fmt"

	"github.com/JosephLai241/shift/utils"
	_ "github.com/mattn/go-sqlite3"
)

// Execute a query within the SQLite instance.
func ExecuteQuery(database *sql.DB, query string) {
	statement, err := database.Prepare(query)
	utils.CheckError(fmt.Sprintf("An error occurred when running the SQL command: \n%s\n\n", query), err)
	statement.Exec()
}
