// Defining the `version` command.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print ASCII art title and the version number",
	Long:  "You should not need more help with this.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
     _   _ ___ _   
 ___| |_|_|  _| |_ 
|_ -|   | |  _|  _|
|___|_|_|_|_| |_|   v1.0.0

`)
	},
}

// Add the version flag to the base command.
func init() {
	rootCmd.AddCommand(versionCmd)
}
