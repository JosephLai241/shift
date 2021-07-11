// Defining the `out` flag.

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/JosephLai241/shift/modify"
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
		fmt.Print(`
         _   
 ___ _ _| |_ 
| . | | |  _|
|___|___|_|

`)

		if status, err := modify.CheckStatus(); !status && err != nil {
			utils.Red.Println("`shift` has not been run.")
			utils.Red.Println("Please initialize the program by recording a shift.")
		} else if !status && err == nil {
			utils.Yellow.Println("`shift` is currently inactive.")
			utils.Yellow.Println("Please clock-in.")
			fmt.Println("")
		} else {
			utils.Green.Print("CLOCKED OUT\n\n")
			currentTime := time.Now().Format("01-02-2006 15:04:05 Mon")
			modify.DisplayStatus()

			message, _ := cmd.Flags().GetString("message")

			ss := modify.ShiftStatus{
				Type:    "OUT",
				Status:  "READY",
				Time:    currentTime,
				Message: message,
			}
			ss.SetStatus()

			shiftData := modify.ShiftData{
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
