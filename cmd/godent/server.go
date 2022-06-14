package main

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/godent/cmd/godent/action/server"
	"github.com/tyrm/godent/cmd/godent/flag"
	"github.com/tyrm/godent/internal/config"
)

// serverCommands returns the 'server' subcommand.
func serverCommands() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "run a godent server",
	}

	serverStartCmd := &cobra.Command{
		Use:   "start",
		Short: "start the godent server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), server.Start)
		},
	}

	flag.Server(serverStartCmd, config.Defaults)

	serverCmd.AddCommand(serverStartCmd)

	return serverCmd
}
