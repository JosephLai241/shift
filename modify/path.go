// Get filepaths used throughout this program.

package modify

import (
	"os"

	"github.com/JosephLai241/shift/utils"
)

// Get the current working directory.
func GetCWD() string {
	cwd, err := os.Getwd()
	utils.CheckError("Could not get the current working directory", err)

	return cwd
}
