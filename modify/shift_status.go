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
var dotfileName = fmt.Sprintf("%s/.%s", cwd, "shiftstatus")

// Check the current shift status.
// Return `false` if `.shiftstatus` dotfile exists or `STATUS` is currently `ACTIVE`.
func CheckStatus() (bool, error) {
	if dotfile, err := os.Open(dotfileName); err != nil {
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
	dotfile, err := os.Open(dotfileName)
	CheckError("Could not open .shiftstatus dotfile", err)
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

// Shift status struct.
type ShiftStatus struct {
	Company string
	Status  string
	Time    string
}

// Set the shift status in the `.shiftstatus` dotfile.
func (ss *ShiftStatus) SetStatus() {
	dotfile, createErr := os.Create(dotfileName)
	CheckError("Could not create .shiftstatus dotfile", createErr)

	status := fmt.Sprintf("STATUS=%s\nCLOCK_IN_TIME=%s", ss.Status, ss.Time)
	if len(ss.Company) > 1 {
		status = status + fmt.Sprintf("\nCOMPANY=%s", ss.Company)
	}

	_, writeErr := dotfile.WriteString(status)
	CheckError("Could not write to .shiftstatus dotfile", writeErr)

	defer dotfile.Close()
}
