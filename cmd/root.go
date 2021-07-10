// Defining the root command.

package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shift",
	Short: "A command-line application for tracking shift (clock-in, clock-out, duration, etc.) data",
	Long: `
     _   _ ___ _   
 ___| |_|_|  _| |_ 
|_ -|   | |  _|  _|
|___|_|_|_|_| |_|

shift is a command-line application designed for contractors/remote workers
who need to keep track of their own shift data. shift can track:

* clock-in time
* clock-out time
* shift duration
* any messages associated with a clock-in or clock-out command call
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// Check if there is an error. Panic if an error is not `nil`.
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
		panic(err)
	}
}

// Initialize clock-in data.
type ShiftData struct {
	day     string // Clock-in day
	time    string // Clock-in time
	message string // Complimentary message
	company string // Complimentary company name
}

// Initialize the directories in which the timesheet will be stored as well as a
// blank timesheet.
func (shiftData *ShiftData) initializeTimesheet() *os.File {
	cwd, err := os.Getwd()
	checkError("Could not get the current working directory", err)

	timesheetDirectory := fmt.Sprintf("%s/shift_timesheets/%s", cwd, time.Now().Format("2006"))
	os.MkdirAll(timesheetDirectory, os.ModePerm)

	currentMonthYear := fmt.Sprintf("%s.csv", time.Now().Format("January"))
	file, err := os.Create(fmt.Sprintf("%s/%s", timesheetDirectory, currentMonthYear))
	checkError("Could not create CSV file", err)
	defer file.Close()

	return file
}

// Write shift data to the `CURRENT_MONTH.csv` file.
func (shiftData ShiftData) writeCSV(file *os.File) {

}

// Write In struct data to CSV.
func (shiftData *ShiftData) RecordClockIn() {
	file := shiftData.initializeTimesheet()
	shiftData.writeCSV(file)

}

// func init() {
// 	cobra.OnInitialize(initConfig)
// 	cobra.OnInitialize(onInit)

// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.
// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shift.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// func onInit() {
// 	fmt.Println(`
//      _   _ ___ _
//  ___| |_|_|  _| |_
// |_ -|   | |  _|  _|
// |___|_|_|_|_| |_|`)
// }

// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".shift" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".shift")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }
