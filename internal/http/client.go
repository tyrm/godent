package http

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"net/http"
)

func NewClient() *http.Client {
	client := http.DefaultClient

	client.Transport = &transport{
		userAgent: fmt.Sprintf("Go-http-client/2.0 (%s/%s; +https://%s/)",
			viper.GetString(config.Keys.ApplicationName),
			viper.GetString(config.Keys.SoftwareVersion),
			viper.GetString(config.Keys.ExternalHostname),
		),
	}

	return client
}

type transport struct {
	userAgent string
}

// RoundTrip executes the default http.Transport with expected http User-Agent.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)

	return http.DefaultTransport.RoundTrip(req)
}
