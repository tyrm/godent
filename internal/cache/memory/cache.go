package memory

import (
	"context"
	"github.com/jellydator/ttlcache/v3"
	"github.com/tyrm/godent/internal/cache"
)

// New creates a new memory cache.
func New(ctx context.Context) (*Cache, error) {
	homeServer := ttlcache.New[string, string](
		ttlcache.WithDisableTouchOnHit[string, string](),
	)
	go homeServer.Start()

	return &Cache{
		homeServer: homeServer,
	}, nil
}

type Cache struct {
	homeServer *ttlcache.Cache[string, string]
}

var _ cache.Cache = (*Cache)(nil)
