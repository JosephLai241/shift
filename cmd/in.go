// Defining the `in` command.

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/JosephLai241/shift/modify"
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

* [-m MESSAGE] - include a message when clocking in
	`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.White.Print(`
 _     
|_|___ 
| |   |
|_|_|_|

`)

		if status, err := modify.CheckStatus(); !status || err != nil {
			utils.Green.Print("CLOCKED IN\n\n")
			currentTime := time.Now().Format("01-02-2006 15:04:05 Mon")
			fmt.Println("Time:", utils.WhiteSprint(currentTime))

			message, _ := cmd.Flags().GetString("message")

			ss := modify.ShiftStatus{
				Type:    "IN",
				Status:  "ACTIVE",
				Time:    currentTime,
				Message: message,
			}
			ss.SetStatus()

			shiftData := modify.ShiftData{
				Type:    "IN",
				Date:    strings.Split(currentTime, " ")[0],
				Day:     time.Now().Format("Monday"),
				Time:    strings.Split(currentTime, " ")[1],
				Message: message,
			}
			shiftData.RecordShift()
		} else {
			utils.Yellow.Print("ALREADY CLOCKED IN\n\n")
			modify.DisplayStatus()
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
