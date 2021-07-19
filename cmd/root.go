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

Author: Joseph Lai
GitHub: https://github.com/JosephLai241/shift

shift is a command-line application designed for contractors/remote workers
who need to keep track of their own working hours. Or for anything else you want 
to track. Inspired by Luke Schenk's Python CLI tool 'clck'.

This program performs CRUD operations on your local machine for the following:

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

Almost all of the commands included in this program contain additional, optional
flags that provide granular control over its operations. I strongly recommend looking
at the help menu for each command to fully take advantage of the features included
for each. You can do so by running 'shift help [COMMAND_NAME]'.

This program allows you to configure how you want to save your recorded shifts.
There are two available options:

- timesheet (CSV spreadsheets)
- database  (relational SQLite database)

Timesheet is the default option; however, you can configure shift to record shifts
to a SQLite database instead by using the 'storage' command. See the help menu for
the command to learn more about how to do so.

If shift is configured to record shifts in timesheets, the directory 'shifts' is
created in your current working directory. This directory contains a sub-directory
labeled with the current year. CSV files labeled with the current month are created
within the year directory, which contain shift data. This is an example of the 'shifts'
directory if you ran shift sometime during July 2021:

shifts/
└── 2021
    └── July.csv

If shift is configured to record shifts in the database instead, 'shifts.db' is
created in your current working directory. shift then creates the main 'YEAR'
table, which holds the current year. The entry then points to a 'Y_CURRENT_YEAR'
table containing the months in which you ran shift. Finally, the months point to a
'M_CURRENT_MONTH' table containing shift data. This is an example of the relationships
within 'shifts.db' if you ran shift sometime during July 2021:

shifts.db
└── TABLE 'YEAR'
    └── TABLE 'Y_2021'
        └── TABLE 'M_July'
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// Initialize the command-line interface.
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
