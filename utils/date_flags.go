// Operations pertaining to the date-related flags (`-d`, `-m`, `-y`).

package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Check whether the input date is valid.
func checkDate(splitDate []string) (bool, string) {
	validMonths := map[string]int{
		"01": 31,
		"02": 28,
		"03": 31,
		"04": 30,
		"05": 31,
		"06": 30,
		"07": 31,
		"08": 30,
		"09": 31,
		"10": 30,
		"11": 31,
		"12": 30,
	}

	if len(splitDate[0]) > 2 || len(splitDate[1]) > 2 || len(splitDate[2]) > 4 {
		return false, "Invalid date. Accepted formats: MM-DD-YYYY, MM/DD/YYYY"
	}

	inputDay, dayErr := strconv.Atoi(splitDate[1])
	if dayErr != nil {
		return false, "Invalid day."
	}
	inputYear, yearErr := strconv.Atoi(splitDate[2])
	if yearErr != nil {
		return false, "Invalid year."
	}
	currentYear, _ := strconv.Atoi(CurrentYear)

	if days, ok := validMonths[splitDate[0]]; !ok {
		return false, "Invalid month."
	} else if ok && inputDay > days {
		return false, "Invalid day."
	} else if ok && inputYear < currentYear {
		return false, "Invalid year."
	}

	return true, "valid"
}

// Check whether the input day of the week is valid.
func checkDays(searchParam string) (bool, string) {
	validDays := map[string]struct{}{
		"monday":    {},
		"tuesday":   {},
		"wednesday": {},
		"thursday":  {},
		"friday":    {},
		"saturday":  {},
		"sunday":    {},
	}

	if _, ok := validDays[strings.ToLower(searchParam)]; !ok {
		return false, "Invalid day of the week."
	}

	return true, "valid"
}

// Call checkDate() and checkDays() in one helper function.
func checkDFlag(searchParam string) (bool, string, string) {
	var newUserInput string
	if strings.Contains(searchParam, "-") {
		splitDate := strings.Split(searchParam, "-")
		if isValid, err := checkDate(splitDate); !isValid {
			return false, searchParam, err
		}
		newUserInput = searchParam
	} else {
		if isValid, err := checkDays(searchParam); !isValid {
			return false, searchParam, err
		}
		newUserInput = strings.Title(searchParam)
	}

	return true, newUserInput, "valid"
}

// Check whether the month is valid.
func checkMonth(searchMonth string) (bool, string, string) {
	validMonths := map[string]string{
		"January":   "01",
		"February":  "02",
		"March":     "03",
		"April":     "04",
		"May":       "05",
		"June":      "06",
		"July":      "07",
		"August":    "08",
		"September": "09",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}

	monthNum, ok := validMonths[searchMonth]
	if !ok {
		return false, "", "Invalid month."
	}

	return true, monthNum, "valid"
}

// Amend the `dayOrDate` parameter if applicable.
func amendDayOrDate(dayOrDate *string) {
	if strings.Contains(*dayOrDate, "/") {
		*dayOrDate = strings.ReplaceAll(*dayOrDate, "/", "-")
	}

	if isValid, fixedDayOrDate, response := checkDFlag(*dayOrDate); !isValid && response != "valid" {
		CheckError("`-d` flag error", errors.New(response))
	} else {
		*dayOrDate = fixedDayOrDate
	}
}

// Update a section of the date within the `dayOrDate` string.
func updateDFlagSection(dayOrDate *string, index int, newString string) {
	splitDFlag := strings.Split(*dayOrDate, "-")
	splitDFlag[index] = newString
	*dayOrDate = strings.Join(splitDFlag, "-")
}

// Amend the `dayOrDate` parameter if the `-m` flag is provided.
func amendMonth(dayOrDate *string, month string) {
	if month != CurrentMonth {
		if isValid, monthNum, response := checkMonth(month); !isValid {
			CheckError("`-m` flag error", errors.New(response))
		} else {
			if strings.Contains(*dayOrDate, "-") {
				updateDFlagSection(dayOrDate, 0, monthNum)
			}
		}
	}
}

// Amend the `dayOrDate` parameter if the `-m` flag is provided.
func amendYear(dayOrDate *string, year string) {
	if year != CurrentYear {
		if strings.Contains(*dayOrDate, "-") {
			updateDFlagSection(dayOrDate, 2, year)
		}
	}
}

// Extract the month and year from the `dayOrDate` flag.
func extractByDate(dayOrDate string) (string, string) {
	splitDate := strings.Split(dayOrDate, "-")
	months := map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}

	month := months[splitDate[0]]
	year := splitDate[2]

	return month, year
}

// Get the `dayOrDate`, `month`, and `year` parameters from flag input.
func FormatFlags(cmd *cobra.Command) (string, string, string) {
	dayOrDate, _ := cmd.Flags().GetString("dayordate")
	month, _ := cmd.Flags().GetString("month")
	month = strings.Title(month)
	year, _ := cmd.Flags().GetString("year")

	amendDayOrDate(&dayOrDate)
	amendMonth(&dayOrDate, month)
	amendYear(&dayOrDate, year)

	if strings.Contains(dayOrDate, "-") && dayOrDate != CurrentDate {
		month, year = extractByDate(dayOrDate)
	}

	return dayOrDate, month, year
}

// Get the timesheet based on the values set by date-related flags.
func GetTimesheetByDFlags(month string, modify bool, year string) (*os.File, error) {
	timesheetPath := fmt.Sprintf(
		"%s/shifts/%s/%s.csv",
		GetCWD(),
		year,
		month,
	)

	var timesheet *os.File
	var err error
	if modify {
		timesheet, err = os.OpenFile(timesheetPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	} else {
		timesheet, err = os.OpenFile(timesheetPath, os.O_RDWR, 0755)
	}

	return timesheet, err
}
