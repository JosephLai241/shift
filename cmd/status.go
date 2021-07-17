// Defining the `status` command.

package cmd

import (
	"fmt"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command.
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the current status of your shift",
	Long: `
     _       _           
 ___| |_ ___| |_ _ _ ___ 
|_ -|  _| .'|  _| | |_ -|
|___|_| |__,|_| |___|___|	

Use this command to check the current status of your shift.

The clock-in time, message, and shift duration is
displayed when shift is currently active.

If shift is currently inactive, a warning message as well
as the previous clock-out time and message is displayed.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.StatusArt)

		if status, err := models.CheckStatus(); !status && err != nil {
			utils.BoldRed.Println("`shift` has not been run.")
			utils.BoldRed.Println("Please initialize the program by recording a shift.")
			fmt.Println("")
		} else if !status {
			utils.BoldYellow.Print("`shift` is currently inactive. Please clock-in.\n\n")
			utils.BoldWhite.Print("Displaying last clock-out information.\n\n")
			models.DisplayStatus(true)
		} else {
			models.DisplayStatus(true)
		}
	},
}

// Add the `status` command and its sub-flags to the base command.
func init() {
	rootCmd.AddCommand(statusCmd)
}
