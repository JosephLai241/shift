// Defining the `affirmation` command.

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
)

// affirmationCmd represents the affirmation command.
var affirmationCmd = &cobra.Command{
	Use:   "affirmation",
	Short: "Get an affirmation",
	Long: `
     ___ ___ _               _   _ 
 ___|  _|  _|_|___ _____ ___| |_|_|___ ___
| .'|  _|  _| |  _|     | .'|  _| | . |   |
|__,|_| |_| |_|_| |_|_|_|__,|_| |_|___|_|_|

Work can be a pain in the ass and everyone needs some support. 
Use this command to display an affirmation.

This command requires an internet connection. If you are not
connected to the internet... sorry, buddy.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("https://www.affirmations.dev/")
		if err != nil {
			utils.CheckError(
				"Sorry, buddy. Could not get an affirmation for you",
				errors.New("you need an internet connection for this"),
			)
		}
		defer response.Body.Close()

		if response.Status != "200 OK" {
			utils.CheckError(
				"Sorry, buddy. Could not get an affirmation for you",
				fmt.Errorf("HTTP Response: %s", response.Status),
			)
		} else {
			var responseMap map[string]string

			err := json.NewDecoder(response.Body).Decode(&responseMap)
			utils.CheckError(
				"Sorry, buddy. An error occurred when reading the response body",
				err,
			)

			if value, ok := responseMap["affirmation"]; ok {
				utils.BoldGreen.Printf("\n%s.\n", value)
				fmt.Println("")
			} else {
				utils.CheckError(
					"Sorry, buddy. Could not get an affirmation for you",
					errors.New("`affirmation` key does not exist in the response body"),
				)
			}
		}
	},
}

// Add the `affirmation` flag to the base command.
func init() {
	rootCmd.AddCommand(affirmationCmd)
}
