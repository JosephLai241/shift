// Defining the `version` command.

package cmd

import (
	"fmt"

	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print ASCII art title and the version number",
	Long:  "You should not need more help with this.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.VersionArt)
	},
}

// Add the `version` flag to the base command.
func init() {
	rootCmd.AddCommand(versionCmd)
}
