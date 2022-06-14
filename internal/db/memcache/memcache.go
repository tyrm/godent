package memcache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/tyrm/godent/internal/db"
	"time"
)

const (
	defaultShards = 32
	defaultHardMaxCacheSize = 8192

	utint64size = 8

	countLifeWindow         = 1 * time.Minute
	countCleanWindow        = 1 * time.Minute
	countMaxEntriesInWindow = 1000
)

// MemCache is an in memory caching middleware for our db interface.
type MemCache struct {
	db db.DB

	count *bigcache.BigCache

	allCaches []*bigcache.BigCache
}

// New creates a new in memory cache.
func New(_ context.Context, d db.DB) (*MemCache, error) {
	count, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             defaultShards,
		LifeWindow:         countLifeWindow,
		CleanWindow:        countCleanWindow,
		MaxEntriesInWindow: countMaxEntriesInWindow,
		MaxEntrySize:       utint64size,
		Verbose:            true,
		HardMaxCacheSize:   defaultHardMaxCacheSize,
	})
	if err != nil {
		return nil, err
	}

	return &MemCache{
		db:      d,

		count: count,

		allCaches: []*bigcache.BigCache{
			count,
		},
	}, nil
}

// Close is a pass through.
func (c *MemCache) Close(ctx context.Context) db.Error {
	for _, cache := range c.allCaches {
		_ = cache.Close()
	}

	return c.db.Close(ctx)
}

// Create is a pass through.
func (c *MemCache) Create(ctx context.Context, i interface{}) db.Error {
	return c.db.Create(ctx, i)
}

// DoMigration is a pass through.
func (c *MemCache) DoMigration(ctx context.Context) db.Error {
	return c.db.DoMigration(ctx)
}

// LoadTestData is a pass through.
func (c *MemCache) LoadTestData(ctx context.Context) db.Error {
	return c.db.LoadTestData(ctx)
}

// ReadByID is a pass through.
func (c *MemCache) ReadByID(ctx context.Context, id int64, i interface{}) db.Error {
	return c.db.ReadByID(ctx, id, i)
}

// ResetCache clears all the caches.
func (c *MemCache) ResetCache(ctx context.Context) db.Error {
	for _, cache := range c.allCaches {
		_ = cache.Reset()
	}

	return c.db.ResetCache(ctx)
}

// Update is a pass through.
func (c *MemCache) Update(ctx context.Context, i interface{}) db.Error {
	return c.db.Update(ctx, i)
}
