// Modify the `.shiftstatus` dotfile.

package models

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JosephLai241/shift/utils"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var cwd = utils.GetCWD()
var DotfileName = fmt.Sprintf("%s/.%s", cwd, "shiftstatus")

// Format the string that will be written to the `.shiftstatus` dotfile.
func formatStatus(io string, ss ShiftStatus) string {
	status := fmt.Sprintf("STATUS=%s\n%s Time=%s\n", ss.Status, io, ss.Time)
	if len(ss.Message) > 1 {
		status += fmt.Sprintf("%s Message=%s\n", io, ss.Message)
	}
	status += fmt.Sprintf("storage-type=%s\n", viper.GetString("storage-type"))

	return status
}

// Shift status struct.
type ShiftStatus struct {
	Type    string // "IN" or "OUT" (clock-in/out)
	Status  string // "ACTIVE" or "READY"
	Time    string // Clock-in or clock-out time
	Message string // Clock-in or clock-out messsage
}

// Set the shift status data in `.shiftstatus`.
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

// Read the `.shiftstatus` dotfile and convert the data into a map.
func readDotfile() (map[string]string, error) {
	dotfileMap := make(map[string]string)

	dotfile, err := os.Open(DotfileName)
	if err != nil {
		return nil, errors.New("'.shiftstatus' does not exist")
	} else {
		scanner := bufio.NewScanner(dotfile)
		for scanner.Scan() {
			splitString := strings.Split(scanner.Text(), "=")
			dotfileMap[splitString[0]] = splitString[1]
		}
	}
	dotfile.Close()

	return dotfileMap, nil
}

// Check the current shift status.
// Return `false` if `.shiftstatus` does not exist or if `STATUS` is not `ACTIVE`.
func CheckStatus() (bool, error) {
	dotfileMap, err := readDotfile()
	if err != nil {
		return false, err
	}
	if status := dotfileMap["STATUS"]; status != "ACTIVE" {
		return false, nil
	}

	return true, nil
}

// Get the current `storage-type` value in `.shiftstatus`.
func GetCurrentStorageType() string {
	dotfileMap, err := readDotfile()
	utils.CheckError("Error reading '.shiftstatus'", err)

	return dotfileMap["storage-type"]
}

// Display the current shift status set in `.shiftstatus`.
func DisplayStatus(displayState bool) {
	dotfileMap, err := readDotfile()
	utils.CheckError("Error reading '.shiftstatus'", err)

	printFields := map[string]struct{}{
		"Clock-in Time":     {},
		"Clock-in Message":  {},
		"Clock-out Time":    {},
		"Clock-out Message": {},
	}

	if status := dotfileMap["STATUS"]; status == "ACTIVE" && displayState {
		utils.BoldGreen.Add(color.Italic).Printf("%s\n\n", status)
	}

	var duration string
	for key, value := range dotfileMap {
		if _, ok := printFields[key]; ok {
			if key == "Clock-in Time" {
				fmt.Printf("%s: %s\n", key, value)
				start, _ := time.ParseInLocation("01-02-2006 15:04:05 Monday", value, time.Now().Location())
				durationObject := time.Since(start)
				duration = time.Time{}.Add(durationObject).Format("15:04:05")
			} else {
				fmt.Printf("%s: %s\n", key, value)
			}
		}
	}
	if len(duration) > 1 {
		utils.BoldBlue.Printf("\nDuration: %s\n", duration)
	}
	fmt.Println("")
}
