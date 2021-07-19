// Defining the `storage` command.

package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// storageCmd represents the storage command.
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Display or configure the current storage type",
	Long: `
     _                       
 ___| |_ ___ ___ ___ ___ ___ 
|_ -|  _| . |  _| .'| . | -_|
|___|_| |___|_| |__,|_  |___|
                    |___|

Use this command to display or configure the storage method.

When used without the optional command 'set', the current
storage method will be displayed.

You can configure the new storage method by including the 'set'
command, followed by the new method. Accepted storage methods are:

- timesheet
- database
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.StorageArt)

		utils.BoldWhite.Printf("Current storage method: %s\n\n", utils.BoldGreen.Sprint(viper.GetString("storage-type")))
	},
}

// setCmd represents the set sub-command.
var setCmd = &cobra.Command{
	Use:   "set (timesheet|database)",
	Short: "Set a new storage method",
	Long: `
Set a new storage method.

Accepted values are:

- timesheet
- database
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.StorageArt)

		newMethod := strings.ToLower(args[0])

		acceptedValues := map[string]struct{}{
			"timesheet": {},
			"database":  {},
		}

		status, _ := models.CheckStatus()
		currentConfig := viper.GetString("storage-type")

		if _, ok := acceptedValues[newMethod]; !ok {
			utils.CheckError(
				"Invalid value",
				errors.New("accepted values are: timesheet, database"),
			)
		} else if ok && status {
			utils.CheckError(
				"`shift` is currently active",
				errors.New("cannot change storage type while recording a shift"),
			)
		} else if ok && newMethod == currentConfig {
			utils.CheckError(fmt.Sprintf(
				"'%s' is already the current storage configuration", currentConfig),
				errors.New("nothing done"),
			)
		} else {
			utils.BoldWhite.Printf("Current storage method: %s\n", utils.BoldRed.Sprint(currentConfig))
			utils.BoldWhite.Printf("New storage method: %s\n\n", utils.BoldGreen.Sprint(newMethod))

			configFile := fmt.Sprintf("%s/%s", utils.GetCWD(), ".shiftconfig.yml")
			viper.SetConfigFile(configFile)
			viper.Set("storage-type", newMethod)

			viper.WriteConfigAs(configFile)
		}
	},
}

// Add the `storage` command and its sub-commands to the base command.
func init() {
	rootCmd.AddCommand(storageCmd)

	storageCmd.AddCommand(setCmd)
}
