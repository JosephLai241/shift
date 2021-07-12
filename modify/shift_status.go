// Modify the `.shiftstatus` dotfile.

package modify

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JosephLai241/shift/utils"
	"github.com/fatih/color"
)

var cwd = utils.GetCWD()
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
	var duration string
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), "=")
		if splitString[0] == "STATUS" {
			if splitString[1] == "ACTIVE" {
				utils.BoldGreen.Add(color.Italic).Printf("%s: %s\n\n", splitString[0], splitString[1])
			} else {
				fmt.Printf("%s: %s\n\n", splitString[0], splitString[1])
			}
		} else {
			fmt.Printf("%s: %s\n", splitString[0], splitString[1])
			if splitString[0] == "Clock-in Time" {
				start, _ := time.Parse("01-02-2006 15:04:05 Mon", splitString[1])
				duration = time.Since(start).String()
			}
		}
	}
	if len(duration) > 1 {
		utils.BoldWhite.Printf("\nShift Duration: %s\n", duration)
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
