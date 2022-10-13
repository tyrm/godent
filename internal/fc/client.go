package fc

import (
	"context"
	"net/http"

	"github.com/jellydator/ttlcache/v3"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func New(httpClient *http.Client) *Client {
	c := Client{
		http:   httpClient,
		tracer: otel.Tracer("internal/fc"),
	}

	c.homeServerCache = ttlcache.New[string, string](
		ttlcache.WithDisableTouchOnHit[string, string](),
	)
	go c.homeServerCache.Start()

	return &c
}

type Client struct {
	http   *http.Client
	tracer trace.Tracer

	// caches
	homeServerCache *ttlcache.Cache[string, string]
}

func (c *Client) get(ctx context.Context, u string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	return c.http.Do(req)
}
