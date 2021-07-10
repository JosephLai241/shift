// Defining the `out` flag.

package cmd

import (
	"fmt"

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

* '-m MESSAGE' - include a message when clocking out
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
         _   
 ___ _ _| |_ 
| . | | |  _|
|___|___|_|

`)
		if message, _ := cmd.Flags().GetString("message"); message != "Clocked out" {
			// fmt.Printf("Clock-in message: %s\n\n", message)
			addClockOutMessage(message)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// outCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// outCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addClockOutMessage(message string) {
	fmt.Printf("Clock-out message: %s\n\n", message)
}
