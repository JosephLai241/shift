// Write shift data to the `CURRENT_MONTH.csv` timesheet.

package modify

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/JosephLai241/shift/utils"
)

// Initialize the directories in which the timesheet will be stored.
// Returns a string denoting the path to the timesheet directory.
func InitializeDirectories() string {
	cwd := utils.GetCWD()
	timesheetDirectory := fmt.Sprintf("%s/shifts/%s", cwd, utils.CurrentYear)
	os.MkdirAll(timesheetDirectory, os.ModePerm)

	return timesheetDirectory
}

// Create the path to the timesheet.
func GetTimesheetPath(timesheetDirectory string) string {
	currentMonthYear := fmt.Sprintf("%s.csv", utils.CurrentMonth)
	timesheetPath := fmt.Sprintf("%s/%s", timesheetDirectory, currentMonthYear)

	return timesheetPath
}

// If the file does not already exist, create a new timesheet with column headers.
// Returns *os.File for writing.
func InitializeTimesheet(timesheetPath string) *os.File {
	var file *os.File
	if _, err := os.Stat(timesheetPath); err == nil {
		file, _ = os.OpenFile(timesheetPath, os.O_RDWR, 0755)
		utils.CheckError("Could not open timesheet", err)
	} else if os.IsNotExist(err) {
		file, err = os.Create(timesheetPath)
		utils.CheckError("Could not create timesheet", err)
	}

	return file
}

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
