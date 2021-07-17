//

package cmd

import (
	"fmt"
	"time"

	"github.com/JosephLai241/shift/modify"
	"github.com/JosephLai241/shift/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a recorded shift",
	Long: `
   _     _     _       
 _| |___| |___| |_ ___ 
| . | -_| | -_|  _| -_|
|___|___|_|___|_| |___|

Use this command to delete a recorded shift. This command 
is fairly versatile - you can search for records based on 
a day of the week or date, month, and/or year.

Using delete without additional commands or flags will display a
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
		fmt.Println(utils.DeleteArt)

		dayOrDate, month, year := amendFlags(cmd)
		fmt.Println(dayOrDate)
		fmt.Println(year)

		modify.CRUD(
			func() { deleteShiftTimesheet(dayOrDate, month, year) },
			func() { deleteShiftDatabase() },
		)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP(
		"dayordate", "d",
		time.Now().Format("01-02-2006"),
		"Search records on a day of the week or date",
	)
	deleteCmd.Flags().StringP(
		"month", "m",
		time.Now().Format("January"),
		"Search records in a month",
	)
	deleteCmd.Flags().StringP(
		"year", "y",
		time.Now().Format("2006"),
		"Search records in a year",
	)
}

// Delete the selected shift from the timesheet.
func deleteShiftTimesheet(dayOrDate string, month string, year string) {
	fmt.Println("deleteShiftTimesheet() called!")
}

// Delete the selected shift from the database.
func deleteShiftDatabase() {
	fmt.Println("deleteShiftDB() called!")
}
