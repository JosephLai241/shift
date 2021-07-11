// Error handling used throughout this application.

package utils

import "log"

// Check if there is an error. Panic if an error is not `nil`.
func CheckError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
		panic(err)
	}
}
