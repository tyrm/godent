package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/godent/internal/config"
)

// Database adds all flags for running the database command.
func Database(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.DBAddress, values.DBAddress, usage.DBAddress)
	cmd.PersistentFlags().Int(config.Keys.DBPort, values.DBPort, usage.DBPort)
	cmd.PersistentFlags().String(config.Keys.DBUser, values.DBUser, usage.DBUser)
	cmd.PersistentFlags().String(config.Keys.DBPassword, values.DBPassword, usage.DBPassword)
	cmd.PersistentFlags().String(config.Keys.DBDatabase, values.DBDatabase, usage.DBDatabase)
	cmd.PersistentFlags().String(config.Keys.DBTLSMode, values.DBTLSMode, usage.DBTLSMode)
	cmd.PersistentFlags().String(config.Keys.DBTLSCACert, values.DBTLSCACert, usage.DBTLSCACert)
}

// DatabaseMigrate adds all flags for running the database migrate command.
func DatabaseMigrate(cmd *cobra.Command, values config.Values) {
}
