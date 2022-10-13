package http

import (
	"net/http"
	"time"

	"github.com/tyrm/godent/internal/util"
)

func NewRetryClient() *http.Client {
	client := http.DefaultClient

	client.Transport = &retryTransport{
		backoffTimes: []time.Duration{
			500 * time.Millisecond,
			1 * time.Second,
			2 * time.Second,
			5 * time.Second,
		},
		userAgent: getUserAgent(),
	}

	return client
}

type retryTransport struct {
	backoffTimes []time.Duration
	userAgent    string
}

// RoundTrip executes the default http.Transport with expected http User-Agent.
func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)

	var ec util.ErrorCollector
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err == nil {
		return resp, nil
	}
	ec.Append(err)
	for _, backoff := range t.backoffTimes {
		select {
		case <-req.Context().Done():
			ec.Append(req.Context().Err())

			return nil, ec.Error()
		case <-time.After(backoff):
			resp, err = http.DefaultTransport.RoundTrip(req)
			if err == nil {
				return resp, nil
			}
			ec.Append(err)
		}
	}

	return nil, ec.Error()
}
