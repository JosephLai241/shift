//

package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/JosephLai241/shift/modify"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// amendCmd represents the amend command
var amendCmd = &cobra.Command{
	Use:   "amend (in|out)",
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

		userInput := checkInput(args, cmd)
		switch storageType := viper.GetString("storage-type"); storageType {
		case "timesheet":
			amendTimesheet(userInput)
		case "database":
			fmt.Println("DATABASE SELECTED")
		}
	},
}

func init() {
	rootCmd.AddCommand(amendCmd)

	amendCmd.Flags().StringP(
		"dayordate", "d",
		time.Now().Format("01-02-2006"),
		"Narrow your search by the day of the week or by a date",
	)
}

// Check whether the input date is valid.
func checkDate(splitDate []string) (bool, string) {
	validMonths := map[string]struct{ days int }{
		"01": {days: 31},
		"02": {days: 28},
		"03": {days: 31},
		"04": {days: 30},
		"05": {days: 31},
		"06": {days: 30},
		"07": {days: 31},
		"08": {days: 30},
		"09": {days: 31},
		"10": {days: 30},
		"11": {days: 31},
		"12": {days: 30},
	}

	if len(splitDate[0]) > 2 || len(splitDate[1]) > 2 || len(splitDate[2]) > 4 {
		return false, "Invalid date."
	}

	inputDay, dayErr := strconv.Atoi(splitDate[1])
	if dayErr != nil {
		return false, "Invalid day."
	}
	inputYear, yearErr := strconv.Atoi(splitDate[2])
	if yearErr != nil {
		return false, "Invalid year."
	}
	currentYear, _ := strconv.Atoi(time.Now().Format("2006"))

	if dayStruct, ok := validMonths[splitDate[0]]; !ok {
		return false, "Invalid month."
	} else if ok && inputDay > dayStruct.days {
		return false, "Invalid day."
	} else if ok && inputYear < currentYear {
		return false, "Invalid year."
	}

	return true, "valid"
}

// Check whether the input day of the week is valid.
func checkDays(userInput string) (bool, string) {
	validDays := map[string]struct{}{
		"monday":    {},
		"tuesday":   {},
		"wednesday": {},
		"thursday":  {},
		"friday":    {},
		"saturday":  {},
		"sunday":    {},
	}

	if _, ok := validDays[strings.ToLower(userInput)]; !ok {
		return false, "Invalid day of the week."
	}

	return true, "valid"
}

// Call checkDate() and checkDays() in one helper function.
func checkDFlag(userInput string) (bool, string) {
	if strings.Contains(userInput, "/") || strings.Contains(userInput, "-") {
		var splitDate []string
		switch true {
		case strings.Contains(userInput, "/"):
			splitDate = strings.Split(userInput, "/")
		case strings.Contains(userInput, "-"):
			splitDate = strings.Split(userInput, "-")
		}

		if isValid, err := checkDate(splitDate); !isValid {
			return false, err
		}
	} else {
		if isValid, err := checkDays(userInput); !isValid {
			return false, err
		}
	}

	return true, "valid"
}

// Check all input for the `amend` command.
func checkInput(args []string, cmd *cobra.Command) string {
	if len(args) < 1 {
		utils.BoldRed.Println("`amend` requires in or out.")
		fmt.Println("")
		os.Exit(1)
	} else if len(args) < 2 {
		utils.BoldRed.Println("`amend` requires a new message.")
		fmt.Println("")
		os.Exit(1)
	} else {
		utils.BoldBlue.Printf("New message: %s\n", args[len(args)-1])
	}

	userInput, _ := cmd.Flags().GetString("dayordate")
	if len(userInput) > 1 {
		if isValid, response := checkDFlag(userInput); !isValid && response != "valid" {
			utils.BoldRed.Printf("\n%s\n", response)
			fmt.Println("")
			os.Exit(1)
		}
	}

	return userInput
}

// Amend a shift.
func amendTimesheet(userInput string) {
	timesheetDirectory := modify.InitializeDirectories()
	timesheetPath := modify.GetTimesheetPath(timesheetDirectory)
	timesheet, err := os.OpenFile(timesheetPath, os.O_RDWR, 0755)
	utils.CheckError("Unable to open the timesheet", err)

	rows := modify.ReadTimesheet(timesheet)
	fmt.Println(rows)
}
