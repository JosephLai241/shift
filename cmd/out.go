// Defining the `out` flag.

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
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

You can also include these sub-commands:

* [-m MESSAGE] - include a message when clocking out
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.OutArt)

		if status, err := models.CheckStatus(); !status && err != nil {
			utils.BoldRed.Println("`shift` has not been run.")
			utils.BoldRed.Println("Please initialize the program by recording a shift.")
		} else if !status && err == nil {
			utils.BoldYellow.Println("`shift` is currently inactive. Please clock-in.")
			fmt.Println("")
		} else {
			currentTime := time.Now().Format("01-02-2006 15:04:05 Mon")
			models.DisplayStatus()

			message, _ := cmd.Flags().GetString("message")

			ss := models.ShiftStatus{
				Type:    "OUT",
				Status:  "READY",
				Time:    currentTime,
				Message: message,
			}
			ss.SetStatus()

			shiftData := models.ShiftData{
				Type:    "OUT",
				Date:    strings.Split(currentTime, " ")[0],
				Day:     time.Now().Format("Monday"),
				Time:    strings.Split(currentTime, " ")[1],
				Message: message,
			}
			shiftData.RecordShift()
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
