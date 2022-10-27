package fc

import (
	"context"
	"net/http"

	"github.com/tyrm/godent/internal/cache"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func New(cacher cache.Cache, httpClient *http.Client) *Client {
	c := Client{
		cache:  cacher,
		http:   httpClient,
		tracer: otel.Tracer("internal/fc"),
	}

	return &c
}

type Client struct {
	http   *http.Client
	cache  cache.Cache
	tracer trace.Tracer

	// caches
	// homeServerCache *ttlcache.Cache[string, string]
}

func (c *Client) get(ctx context.Context, u string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	return c.http.Do(req)
}
