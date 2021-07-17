// Defining the execution switch - determines where to execute CRUD operations.

package modify

import "github.com/spf13/viper"

// Main CRUD switch execute the timesheet or database function depending on the
// `storage-type` value set in `.shiftconfig.yml`.
func CRUD(timesheetFn func(), databaseFn func()) {
	switch storageType := viper.GetString("storage-type"); storageType {
	case "timesheet":
		timesheetFn()
	case "database":
		databaseFn()
	}
}
