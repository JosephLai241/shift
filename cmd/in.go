// Defining the `in` command.

package cmd

import (
	"fmt"
	"strings"
	"time"

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

* [-m MESSAGE]      - include a message when clocking in
* [-c COMPANY_NAME] - include a company name associated with the clock-in
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
 _     
|_|___ 
| |   |
|_|_|_|

`)

		currentTime := time.Now().Format("01-02-2006 15:04:05 Mon")
		fmt.Println("Current time:", currentTime)

		message, _ := cmd.Flags().GetString("message")
		companyName, _ := cmd.Flags().GetString("company")

		if len(companyName) > 1 {
			fmt.Printf("Company: %s\n", companyName)
		}
		if len(message) > 1 {
			fmt.Printf("Message: %s\n\n", message)
		}

		shiftData := ShiftData{
			day:     strings.Split(currentTime, " ")[2],
			time:    strings.Split(currentTime, " ")[1],
			message: message,
			company: companyName,
		}
		shiftData.RecordClockIn()
		// fmt.Printf("%+v\n", in)
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
	inCmd.PersistentFlags().StringP(
		"company", "c",
		"",
		"Include a complimentary company name associated with the clock-in",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}