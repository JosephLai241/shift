// Error handling used throughout this application.

package utils

import (
	"errors"
	"fmt"
	"log"
)

// Check if there is an error. Panic if an error is not `nil`.
func CheckError(message string, err error) {
	if err != nil {
		fmt.Println(ErrorArt)
		log.Fatal(BoldRed.Sprintf("\n%s: ", message), err)
	}
}

// Format and then display the error message if no matches were found.
func NoMatchesError(dayOrDate string, month string, year string) {
	var errorMessage error
	if dayOrDate != CurrentDate || month != CurrentMonth || year != CurrentYear {
		errorMessage = errors.New("no shifts were found based on your search parameters")
	} else {
		errorMessage = errors.New("no shifts were recorded today")
	}
	CheckError("Error", errorMessage)
}
