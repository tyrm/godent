package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/godent/internal/config"
)

// Server adds all flags for running the server.
func Server(cmd *cobra.Command, values config.Values) {
	Database(cmd, values)
	Redis(cmd, values)

	// server
	cmd.PersistentFlags().String(config.Keys.ExternalHostname, values.ExternalHostname, usage.ExternalHostname)
	cmd.PersistentFlags().String(config.Keys.ServerHTTPBind, values.ServerHTTPBind, usage.ServerHTTPBind)

	// server
	cmd.PersistentFlags().Bool(config.Keys.RequireTermsAgreed, values.RequireTermsAgreed, usage.RequireTermsAgreed)
}
