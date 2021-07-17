// Defining the `out` command.

package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/modify"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// outCmd represents the out command.
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
			runStorageCheck()
			modify.CRUD(
				func() { recordOutTimesheet(message) },
				func() { recordOutDatabase(message) },
			)
		}
	},
}

// Add the `out` command and its sub-flags to the base command.
func init() {
	rootCmd.AddCommand(outCmd)

	outCmd.PersistentFlags().StringP(
		"message", "m",
		"Clocked out",
		"Include a complimentary clock-out message",
	)
}

// Check the storage setting written in `.shiftstatus`.
func runStorageCheck() {
	if storageType := models.CheckStorageType(); storageType != viper.GetString("storage-type") {
		issueString := "The `storage-type` value was changed while you were clocked-in."
		previousString := fmt.Sprintf("\n\nThe previous value was set at: %s\n", utils.BoldWhite.Sprint(storageType))
		currentString := fmt.Sprintf("But the current value is set at: %s\n\n", utils.BoldWhite.Sprint(viper.GetString("storage-type")))
		suggestionString := "Please change the `storage-type` value in '.shiftconfig.yml' back to the previous value before clocking out.\n\n"

		utils.CheckError(
			"`storage-type` error",
			errors.New(fmt.Sprint(
				issueString+previousString+currentString+suggestionString),
			),
		)
	}
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
