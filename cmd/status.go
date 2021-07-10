// Defining the `status` command.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the current status of your shift",
	Long: `
     _       _           
 ___| |_ ___| |_ _ _ ___ 
|_ -|  _| .'|  _| | |_ -|
|___|_| |__,|_| |___|___|	

Use this command to check the current status of your shift.
Clock-in time and shift duration will be displayed. 

The complimentary shift message and company name will also be
displayed, if applicable, by including the [-v] flag
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
     _       _           
 ___| |_ ___| |_ _ _ ___ 
|_ -|  _| .'|  _| | |_ -|
|___|_| |__,|_| |___|___|

`)
		verboseStatus, _ := cmd.Flags().GetBool("verbose")
		if verboseStatus {
			printVerboseDetails()
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().BoolP(
		"verbose", "v",
		false,
		"Print verbose shift status",
	)
}

func printVerboseDetails() {
	fmt.Println("VERBOSE DATA TRIGGERED")
	// shiftData := ShiftData{
	// 	day: day,
	// 	time: time,
	// 	message: message,
	// 	company: company,
	// }
}
