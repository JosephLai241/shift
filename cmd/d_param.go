// Amend the `dayOrDate` string.

package cmd

import (
	"errors"
	"strings"
	"time"

	"github.com/JosephLai241/shift/utils"
)

// Amend the `dayOrDate` parameter if applicable.
func amendDayOrDate(dayOrDate *string) {
	if strings.Contains(*dayOrDate, "/") {
		*dayOrDate = strings.ReplaceAll(*dayOrDate, "/", "-")
	}

	if isValid, fixedDayOrDate, response := checkDFlag(*dayOrDate); !isValid && response != "valid" {
		utils.CheckError("`-d` flag error", errors.New(response))
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
	if month != time.Now().Format("January") {
		month = strings.Title(month)
		if isValid, monthNum, response := checkMonth(month); !isValid {
			utils.CheckError("`-m` flag error", errors.New(response))
		} else {
			if strings.Contains(*dayOrDate, "-") {
				updateDFlagSection(dayOrDate, 0, monthNum)
			}
		}
	}
}

// Amend the `dayOrDate` parameter if the `-m` flag is provided.
func amendYear(dayOrDate *string, year string) {
	if year != time.Now().Format("2006") {
		if strings.Contains(*dayOrDate, "-") {
			updateDFlagSection(dayOrDate, 2, year)
		}
	}
}
