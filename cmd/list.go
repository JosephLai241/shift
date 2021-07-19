// Defining the `list` command.

package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JosephLai241/shift/database"
	"github.com/JosephLai241/shift/timesheet"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
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

You can display shifts on a different day or date by using the '-d'
flag, which accepts a day of the week (ie. Monday) or a date 
(ie. 07-14-2021). The accepted date formats are:

- MM-DD-YYYY
- MM/DD/YYYY

You can display shifts in a different month by using the
'-m' flag, which accepts a month (ie. July). If this is the only
provided flag, a search will be done for the current day within
the provided month.

Finally, you can display shifts in a different year by using
the '-y' flag, which accepts a year (ie. 2021). Like the '-m'
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

		dayOrDate, month, year := utils.FormatFlags(cmd)

		utils.CRUD(
			func() { timesheet.List(dayOrDate, month, subCommand, year) },
			func() { database.List(dayOrDate, month, subCommand, year) },
		)
	},
}

// Add the `list` command and its sub-flags to the base command.
func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP(
		"dayordate", "d",
		utils.CurrentDate,
		"Narrow your search by the day of the week or by a date",
	)
	listCmd.Flags().StringP(
		"month", "m",
		utils.CurrentMonth,
		"List records in a specific month",
	)
	listCmd.Flags().StringP(
		"year", "y",
		utils.CurrentYear,
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
