//

package cmd

import (
	"fmt"

	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `
   _     _     _       
 _| |___| |___| |_ ___ 
| . | -_| | -_|  _| -_|
|___|___|_|___|_| |___|

Use this command to delete a recorded shift. You can
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.DeleteArt)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
