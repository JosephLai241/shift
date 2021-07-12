//

package cmd

import (
	"fmt"

	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
)

// amendCmd represents the amend command
var amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "A brief description of your command",
	Long: `
                     _ 
 ___ _____ ___ ___ _| |
| .'|     | -_|   | . |
|__,|_|_|_|___|_|_|___|

Use this command to amend a recorded shift's clock-in or clock-out message.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.AmendArt)
	},
}

func init() {
	rootCmd.AddCommand(amendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// amendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// amendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
