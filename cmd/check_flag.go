// Storing functions that may be reused to check flags.

package cmd

import (
	"strconv"
	"strings"
	"time"
)

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
