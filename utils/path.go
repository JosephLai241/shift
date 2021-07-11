// Get filepaths used throughout this program.

package utils

import "os"

// Get the current working directory.
func GetCWD() string {
	cwd, err := os.Getwd()
	CheckError("Could not get the current working directory", err)

	return cwd
}
