// Display shifts in an ASCII table.

package views

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Display table with the new changes.
func Display(amendRow [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Date",
		"Day",
		"Clock-In",
		"Clock-In Message",
		"Clock-Out",
		"Clock-Out Message",
		"Shift Duration",
	})

	table.SetHeaderColor(
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
	)

	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(amendRow)
	table.Render()
}

// Display shifts in an ASCII table
func DisplayOptions(rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Shift",
		"Date",
		"Day",
		"Clock-In",
		"Clock-In Message",
		"Clock-Out",
		"Clock-Out Message",
		"Shift Duration",
	})

	table.SetHeaderColor(
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
		tablewriter.Colors{
			tablewriter.Bold,
			tablewriter.FgCyanColor,
		},
	)

	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(rows)
	table.Render()
}
