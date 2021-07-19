// User confirmation.

package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Loop to confirm "Y/N" options.
func ConfirmInput(action string) string {
	validOptions := map[string]int{
		"y": 1,
		"n": 0,
	}

	var input string

	for {
		fmt.Printf("\nConfirm %s? [y/n] ", action)
		fmt.Scanln(&input)

		if _, ok := validOptions[strings.ToLower(input)]; !ok {
			BoldRed.Print("\nInvalid option. Try again.\n")
		} else {
			break
		}
	}

	return strings.ToLower(input)
}

// Check if the selection pointing to a target shift is valid.
func CheckSelection(rowNums []int) int {
	validOptions := make(map[int]struct{})
	for i := range rowNums {
		validOptions[rowNums[i]] = struct{}{}
	}

	var input string
	for {
		fmt.Printf("\nSelect a shift to modify %+v: ", rowNums)
		fmt.Scanln(&input)

		intSelection, _ := strconv.Atoi(input)
		if _, ok := validOptions[intSelection]; !ok {
			BoldRed.Print("\nInvalid option. Try again.\n")
		} else {
			return intSelection
		}
	}
}
