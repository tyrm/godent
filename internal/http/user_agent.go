package http

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
)

func getUserAgent() string {
	return fmt.Sprintf("Go-http-client/2.0 (%s/%s; +https://%s/)",
		viper.GetString(config.Keys.ApplicationName),
		viper.GetString(config.Keys.SoftwareVersion),
		viper.GetString(config.Keys.ExternalHostname),
	)
}
