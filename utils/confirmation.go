// User confirmation.

package utils

import (
	"fmt"
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
