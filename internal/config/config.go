package config

import (
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Init starts config collection.
func Init(flags *pflag.FlagSet) error {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	return viper.BindPFlags(flags)
}

// ReadConfigFile reads the config file from disk if config path is sent.
func ReadConfigFile() error {
	configPath := viper.GetString(Keys.ConfigPath)
	if configPath != "" {
		viper.SetConfigFile(configPath)

		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}

	return nil
}
