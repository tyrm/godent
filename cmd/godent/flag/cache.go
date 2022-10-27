package flag

import (
	"github.com/spf13/cobra"
	"github.com/tyrm/godent/internal/config"
)

func cache(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.CacheStore, values.CacheStore, usage.CacheStore)
}
