// Error handling used throughout this application.

package utils

import (
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
