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
who need to keep track of their own working hours.

This program will perform CRUD operations on your local machine for the following:

- clock-in time
- clock-out time
- shift duration
- any messages associated with a clock-in or clock-out command call

The commands you will likely interact with most are:

- in
- status
- out

There are additional commands that may be very useful to you as your records grow
in size. These commands are:

- amend
- list
- delete

Many of the commands included in this program contain additional, optional flags
that provide granular control over operations such as amending a record's clock-in
or clock-out message, listing stored records, or deleting a record. I strongly
recommend looking at the help menu for each command to fully take advantage of
the features included for each. You can do so by running 'shift help [COMMAND_NAME]'.

By default, this program will create a directory in your current working
directory named 'shifts' and record shift data into CSV files labeled with
the month. You can modify the 'storage-type' variable within '.shiftconfig.yml'
to record shifts in a local SQLite instance named 'shifts.db' (instead of writing
to the default timesheets) by changing the default value from 'timesheet' to 
'database' - these are the only accepted values.
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
