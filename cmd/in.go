// Defining the `in` command.

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
recorded to the timesheet or database.

You can also include a clock-in message by including the '-m' flag
and typing your message in quotes. If you are familiar with using git
from the command line, this is identical to how the 'git commit -m' 
command functions.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.InArt)

		message, _ := cmd.Flags().GetString("message")

		if status, err := models.CheckStatus(); !status || err != nil {
			switch storageType := viper.GetString("storage-type"); storageType {
			case "timesheet":
				recordInTimesheet(message)
			case "database":
				recordInDatabase(message)
			}
		} else {
			utils.BoldYellow.Print("ALREADY CLOCKED IN\n\n")
			models.DisplayStatus(true)
		}
	},
}

// Add the in flag and its subcommands to the base command.
func init() {
	rootCmd.AddCommand(inCmd)

	inCmd.PersistentFlags().StringP(
		"message", "m",
		"Clocked in",
		"Include a complimentary clock-in message",
	)
}

// Record clock-in in the timesheet.
func recordInTimesheet(message string) {
	currentTime := time.Now().Format("01-02-2006 15:04:05 Monday")
	utils.BoldBlue.Println("Clock-in time:", currentTime)
	fmt.Println("")

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
		Day:     strings.Split(currentTime, " ")[2],
		Time:    strings.Split(currentTime, " ")[1],
		Message: message,
	}
	shiftData.RecordShift()
}

// Record clock-in in the database.
func recordInDatabase(message string) {
	fmt.Println("DATABASE SELECTED")
}
