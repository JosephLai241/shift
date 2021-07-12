//

package cmd

import (
	"fmt"

	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `
 _ _     _   
| |_|___| |_ 
| | |_ -|  _|
|_|_|___|_|

Use this command to list all recorded shifts within
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.ListArt)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
