// Defining the root command.

package cmd

import (
	"errors"
	"fmt"

	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "shift",
	Short: "A command-line application for tracking shift (clock-in, clock-out, duration, etc.) data",
	Long: `
     _   _ ___ _   
 ___| |_|_|  _| |_ 
|_ -|   | |  _|  _|
|___|_|_|_|_| |_|

shift is a command-line application designed for contractors/remote workers
who need to keep track of their own shift data. shift will record:

* the current date
* clock-in time
* clock-out time
* shift duration
* any messages associated with a clock-in or clock-out command call
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in the `.shiftconfig.yml` config file.
func initConfig() {
	configFile := fmt.Sprintf("%s/%s", utils.GetCWD(), ".shiftconfig.yml")
	viper.SetConfigFile(configFile)
	viper.SetDefault("storage-type", "timesheet")

	err := viper.ReadInConfig()
	if err != nil {
		viper.SafeWriteConfigAs(configFile)
	}

	envString := viper.GetString("storage-type")

	acceptedValues := map[string]struct{}{
		"timesheet": {},
		"database":  {},
	}
	if _, ok := acceptedValues[envString]; !ok {
		utils.CheckError("`.shiftconfig.yml` error", errors.New(`

The "storage-type" value is invalid. Accepted values are:

- timesheet
- database
		`))
	}
}
