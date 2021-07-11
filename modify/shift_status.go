// Modify the `.shiftstatus` dotfile.

package modify

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/JosephLai241/shift/utils"
)

var cwd = GetCWD()
var DotfileName = fmt.Sprintf("%s/.%s", cwd, "shiftstatus")

// Check the current shift status.
// Return `false` if `.shiftstatus` dotfile exists or `STATUS` is currently `ACTIVE`.
func CheckStatus() (bool, error) {
	if dotfile, err := os.Open(DotfileName); err != nil {
		return false, errors.New("'.shiftstatus' does not exist")
	} else {
		scanner := bufio.NewScanner(dotfile)
		for scanner.Scan() {
			splitString := strings.Split(scanner.Text(), "=")
			if splitString[0] == "STATUS" && splitString[1] == "ACTIVE" {
				return true, nil
			}
		}
		return false, nil
	}
}

// Display the current shift status set in `.shiftstatus`.
func DisplayStatus() {
	dotfile, err := os.Open(DotfileName)
	utils.CheckError("Could not open .shiftstatus dotfile", err)
	defer dotfile.Close()

	scanner := bufio.NewScanner(dotfile)
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), "=")
		if splitString[0] != "STATUS" {
			utils.Blue.Printf("%s: %s\n", splitString[0], splitString[1])
		}
	}
	fmt.Println("")
}

// Format the string that will be written to the `.shiftstatus` dotfile.
func formatStatus(io string, ss ShiftStatus) string {
	status := fmt.Sprintf("STATUS=%s\n%s Time=%s\n", ss.Status, io, ss.Time)
	if len(ss.Message) > 1 {
		status += fmt.Sprintf("%s Message=%s", io, ss.Message)
	}

	return status
}

// Shift status struct.
type ShiftStatus struct {
	Type    string
	Status  string
	Time    string
	Message string
}

// Set the shift status in the `.shiftstatus` dotfile.
func (ss *ShiftStatus) SetStatus() {
	dotfile, createErr := os.Create(DotfileName)
	utils.CheckError("Could not create .shiftstatus dotfile", createErr)

	var io string
	if ss.Type == "IN" {
		io = "Clock-in"
	} else {
		io = "Clock-out"
	}
	status := formatStatus(io, *ss)

	_, writeErr := dotfile.WriteString(status)
	utils.CheckError("Could not write to .shiftstatus dotfile", writeErr)

	dotfile.Close()
}
