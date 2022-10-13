package http

import (
	"net/http"
)

func NewClient() *http.Client {
	client := http.DefaultClient

	client.Transport = &transport{
		userAgent: getUserAgent(),
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
