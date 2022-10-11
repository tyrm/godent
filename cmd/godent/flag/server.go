package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/godent/internal/config"
)

// Server adds all flags for running the server.
func Server(cmd *cobra.Command, values config.Values) {
	Redis(cmd, values)

	// server
	cmd.PersistentFlags().String(config.Keys.ServerHTTPBind, values.ServerHTTPBind, usage.ServerHTTPBind)
}
