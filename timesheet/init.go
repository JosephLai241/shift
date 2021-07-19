// Initialize the timesheet.

package timesheet

import (
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
