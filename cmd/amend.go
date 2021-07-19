// Defining the `amend` command.

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

// amendCmd represents the amend command.
var amendCmd = &cobra.Command{
	Use:   `amend (in|out) "NEW MESSAGE"`,
	Short: "Amend a shift's clock-in or clock-out message",
	Long: `
                     _ 
 ___ _____ ___ ___ _| |
| .'|     | -_|   | . |
|__,|_|_|_|___|_|_|___|

Use this command to amend a recorded shift's clock-in or clock-out
message. This command is fairly versatile - you can search for records 
based on a day of the week or date, month, and/or year.

Using amend without additional commands or flags will display a
table containing shifts recorded for the current day.

There are three optional flags you can use: the '-d', '-m', 
and '-y' flags. These flags denote the target day of the week or date, 
month, and year respectively. The default value for all of these 
flags is the current day of the week/date, month, and year. 
Combine these flags to to do a deep search for a particular 
shift or shifts.

You can search for shifts on a different day or date by using the '-d'
flag, which accepts a day of the week (ie. Monday) or a date 
(ie. 07-14-2021). The accepted date formats are:

- MM-DD-YYYY
- MM/DD/YYYY

You can search for shifts in a different month by using the
'-m' flag, which accepts a month (ie. July). If this is the only
provided flag, a search will be done for the current day within
the provided month.

Finally, you can search for shifts in a different year by using
the '-y' flag, which accepts a year (ie. 2021). Like the '-m'
flag, a search will be done for the current day and month within
the provided year if this is the only provided flag.

You can combine the '-d', '-m', and/or '-y' flags to do a deep
search for a particular shift or shifts.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.AmendArt)

		checkArgs(args)
		args[0] = strings.ToLower(args[0])
		dayOrDate, month, year := utils.FormatFlags(cmd)

		utils.CRUD(
			func() { timesheet.Amend(args, dayOrDate, month, year) },
			func() { database.Amend(args, dayOrDate, month, year) },
		)
	},
}

// Add the `amend` command and its sub-flags to the base command.
func init() {
	rootCmd.AddCommand(amendCmd)

	amendCmd.Flags().StringP(
		"dayordate", "d",
		utils.CurrentDate,
		"Search records on a day of the week or date",
	)
	amendCmd.Flags().StringP(
		"month", "m",
		utils.CurrentMonth,
		"Search records in a month",
	)
	amendCmd.Flags().StringP(
		"year", "y",
		utils.CurrentYear,
		"Search records in a year",
	)
}

// Check all input for the `amend` command.
func checkArgs(args []string) {
	if len(args) < 1 {
		utils.CheckError("Command error", errors.New("`amend` requires in or out"))
	} else if len(args) < 2 {
		utils.CheckError("Command error", errors.New("`amend` requires a new message"))
	} else {
		utils.BoldBlue.Printf("New message: %s\n", args[len(args)-1])
	}
}
