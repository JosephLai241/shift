// Defining the `out` flag.

package cmd

import (
	"fmt"
	"time"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// outCmd represents the out command
var outCmd = &cobra.Command{
	Use:   "out",
	Short: "Clock out",
	Long: `
         _   
 ___ _ _| |_ 
| . | | |  _|
|___|___|_|

Use this command to clock-out. The current time will be
recorded to the timesheet.

You can also include a clock-out message by including the '-m' flag
and typing your message in quotes. If you are familiar with using git
from the command line, this is identical to how the 'git commit -m' 
command functions.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.OutArt)

		message, _ := cmd.Flags().GetString("message")

		if status, err := models.CheckStatus(); !status && err != nil {
			utils.BoldRed.Println("`shift` has not been run.")
			utils.BoldRed.Println("Please initialize the program by recording a shift.")
		} else if !status && err == nil {
			utils.BoldYellow.Println("`shift` is currently inactive. Please clock-in.")
			fmt.Println("")
		} else {
			switch storageType := viper.GetString("storage-type"); storageType {
			case "timesheet":
				recordOutTimesheet(message)
			case "database":
				recordOutDatabase(message)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(outCmd)

	outCmd.PersistentFlags().StringP(
		"message", "m",
		"Clocked out",
		"Include a complimentary clock-out message",
	)
}

// Record clock-out in the timesheet.
func recordOutTimesheet(message string) {
	currentTime := time.Now().Format("01-02-2006 15:04:05 Monday")
	models.DisplayStatus(false)

	ss := models.ShiftStatus{
		Type:    "OUT",
		Status:  "READY",
		Time:    currentTime,
		Message: message,
	}
	ss.SetStatus()

	shiftData := models.ShiftData{
		Type:    "OUT",
		Date:    "",
		Day:     "",
		Time:    currentTime,
		Message: message,
	}
	shiftData.RecordShift()
}

// Record clock-out in the database.
func recordOutDatabase(message string) {
	fmt.Println("DATABASE SELECTED")
}
