// Defining the `in` command.

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
)

// inCmd represents the in command
var inCmd = &cobra.Command{
	Use:   "in",
	Short: "Clock in",
	Long: `
 _     
|_|___ 
| |   |
|_|_|_|

Use this command to clock-in. The current time will be
recorded to the timesheet.

You can also include these sub-commands:

* [-m "YOUR MESSAGE HERE"] - include a message when clocking in
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.InArt)

		if status, err := models.CheckStatus(); !status || err != nil {
			recordShift(cmd)
		} else {
			utils.BoldYellow.Print("ALREADY CLOCKED IN\n\n")
			models.DisplayStatus()
		}
	},
}

// Add the in flag and its subcommands to the base command.
func init() {
	rootCmd.AddCommand(inCmd)

	inCmd.PersistentFlags().StringP(
		"message", "m",
		"",
		"Include a complimentary clock-in message",
	)
}

// Record the shift.
func recordShift(cmd *cobra.Command) {
	currentTime := time.Now().Format("01-02-2006 15:04:05 Mon")
	utils.BoldBlue.Println("Clock-in time:", currentTime)
	fmt.Println("")

	message, _ := cmd.Flags().GetString("message")

	ss := models.ShiftStatus{
		Type:    "IN",
		Status:  "ACTIVE",
		Time:    currentTime,
		Message: message,
	}
	ss.SetStatus()

	shiftData := models.ShiftData{
		Type:    "IN",
		Date:    strings.Split(currentTime, " ")[0],
		Day:     time.Now().Format("Monday"),
		Time:    strings.Split(currentTime, " ")[1],
		Message: message,
	}
	shiftData.RecordShift()
}
