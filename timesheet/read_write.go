// Read or write data to timesheets.

package timesheet

import (
	"encoding/csv"
	"os"

	"github.com/JosephLai241/shift/utils"
)

// Write data to to the timesheet.
func WriteToTimesheet(file *os.File, rows [][]string) {
	writer := csv.NewWriter(file)
	err := writer.WriteAll(rows)
	utils.CheckError("Could not flush writer data after writing data to timesheet", err)

	closeErr := file.Close()
	utils.CheckError("Could not write shift data to timesheet", closeErr)
}

// Read timesheet and extract the data into a nested list of strings.
func ReadTimesheet(file *os.File) [][]string {
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	utils.CheckError("Could not read timesheet", err)

	file.Close()
	return rows
}
