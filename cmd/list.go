//

package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JosephLai241/shift/models"
	"github.com/JosephLai241/shift/modify"
	"github.com/JosephLai241/shift/utils"
	"github.com/JosephLai241/shift/views"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list [all]",
	Short: "Display recorded shifts",
	Long: `
 _ _     _   
| |_|___| |_ 
| | |_ -|  _|
|_|_|___|_|

Use this command to list recorded shifts. This command is
fairly versatile - you can list records based on a day of the
week or date, month, and/or year.

Using list without additional commands or flags will display a
table containing shifts recorded for the current day.

The optional positional argument 'all' following the parent 'list' 
command will display all recorded shifts for the target month. 
You can combine the 'all' argument with the '-m' flag to display all
recorded shifts in a different month. Information for using the 
'-m' flag is provided below.

There are three optional flags you can use: the '-d', '-m', 
and '-y' flags. These flags denote the target day of the week or date, 
month, and year respectively. The default value for all of these 
flags is the current day of the week/date, month, and year. 
Combine these flags to to do a deep search for a particular 
shift or shifts.

You can display shifts for a different day by using the '-d'
flag, which can take in a day of the week (ie. Monday) or a
date (ie. 07-14-2021). The accepted date formats are:

- MM-DD-YYYY
- MM/DD/YYYY

You can display shifts for a different month by using the
'-m' flag, which takes in a month (ie. July). If this is the only
provided flag, a search will be done for the current day within
the provided month.

Finally, you can display shifts for a different year by using
the '-y' flag, which takes in a year (ie. 2021). Like the '-m'
flag, a search will be done for the current day and month within
the provided year if this is the only provided flag.

You can combine the '-d', '-m', and/or '-y' flags to do a deep
search for a particular shift or shifts.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.ListArt)

		var subCommand string
		if len(args) > 0 {
			subCommand = checkOptionalCommand(args)
		}

		dayOrDate, _ := cmd.Flags().GetString("dayordate")
		month, _ := cmd.Flags().GetString("month")
		year, _ := cmd.Flags().GetString("year")

		amendDayOrDate(&dayOrDate)
		amendMonth(&dayOrDate, month)
		amendYear(&dayOrDate, year)

		switch storageType := viper.GetString("storage-type"); storageType {
		case "timesheet":
			timesheetList(dayOrDate, month, subCommand, year)
		case "database":
			fmt.Println("DATABASE SELECTED")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP(
		"dayordate", "d",
		time.Now().Format("01-02-2006"),
		"Narrow your search by the day of the week or by a date",
	)
	listCmd.Flags().StringP(
		"month", "m",
		time.Now().Format("January"),
		"List records in a specific month",
	)
	listCmd.Flags().StringP(
		"year", "y",
		time.Now().Format("2006"),
		"List records in a specific year",
	)
}

// Check whether the optional command is valid.
func checkOptionalCommand(args []string) string {
	validCommands := map[string]struct{}{
		"all": {},
		"":    {},
	}

	if _, ok := validCommands[args[0]]; !ok {
		utils.CheckError(
			"Command error",
			errors.New("invalid command following `list`. Takes `all` or no command"),
		)
	}

	return strings.ToLower(args[0])
}

// Find matches based on search day or date.
func listDMatches(fixedSearchParam string, rows [][]string) {
	_, matches := models.FindMatches(fixedSearchParam, rows)
	if len(matches) == 0 {
		utils.CheckError(
			"Error",
			errors.New("no shifts were found based on your search parameters"),
		)
	}

	for i, row := range matches {
		matches[i] = row[1:]
	}

	views.Display(matches)
}

// Pull and list records from timesheets.
func timesheetList(dayOrDate string, month string, subCommand string, year string) {
	month = strings.Title(month)

	timesheetPath := fmt.Sprintf("%s/shifts/%s/%s.csv", utils.GetCWD(), year, month)
	timesheet, err := os.OpenFile(timesheetPath, os.O_RDWR, 0755)
	if err != nil {
		utils.CheckError(
			fmt.Sprintf("An error occurred when listing shifts recorded in %s %s", month, year),
			errors.New("no shifts were recorded"),
		)
	}

	if rows := modify.ReadTimesheet(timesheet); len(rows) == 0 {
		utils.CheckError(
			"Error",
			errors.New("no shifts were found based on your search parameters"),
		)
	} else {
		if subCommand == "all" {
			utils.BoldBlue.Printf("Displaying all shifts recorded in %s %s.\n\n", month, year)
			views.Display(rows)
		} else {
			if strings.Contains(dayOrDate, "-") {
				utils.BoldBlue.Printf("Displaying shifts recorded on %s.\n\n", dayOrDate)
			} else {
				utils.BoldBlue.Printf("Displaying shifts recorded on a %s within %s %s.\n\n", dayOrDate, month, year)
			}

			listDMatches(dayOrDate, rows)
		}
	}
}
