package main

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/godent/cmd/godent/action/generate"
)

// generateCommands returns the 'generate' subcommand.
func generateCommands() *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "generate config values",
	}

	generateMigrateCmd := &cobra.Command{
		Use:   "signing-key",
		Short: "generate an ed25519 signing key",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), generate.SigningKey)
		},
	}
	generateCmd.AddCommand(generateMigrateCmd)

	return generateCmd
}
