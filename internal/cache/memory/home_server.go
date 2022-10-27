package memory

import (
	"context"
	"time"

	"github.com/tyrm/godent/internal/cache"
)

func (c Cache) DeleteHomeServer(_ context.Context, domain string) error {
	c.homeServer.Delete(domain)

	return nil
}

func (c Cache) GetHomeServer(_ context.Context, domain string) (string, error) {
	cached := c.homeServer.Get(domain)
	if cached == nil {
		return "", cache.ErrMiss
	}

	return cached.Value(), nil
}

func (c Cache) SetHomeServer(ctx context.Context, domain, homeServer string, expiration int) error {
	c.homeServer.Set(domain, homeServer, time.Duration(expiration)*time.Second)

	return nil
}
