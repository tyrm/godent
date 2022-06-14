package redis

import (
	"context"

	redis "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
)

// Client represents a redis client.
type Client struct {
	redis *redis.Client
}

// New creates a new redis client.
func New(ctx context.Context) (*Client, error) {
	l := logger.WithField("func", "New")

	r := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(config.Keys.RedisAddress),
		Password: viper.GetString(config.Keys.RedisPassword),
		DB:       viper.GetInt(config.Keys.RedisDB),
	})

	c := Client{
		redis: r,
	}

	resp := c.redis.Ping(ctx)
	l.Debugf("ping: %s", resp.String())

	return &c, nil
}

// Close closes the redis pool.
func (c *Client) Close(_ context.Context) error {
	return c.redis.Close()
}

// RedisClient returns the redis client.
func (c *Client) RedisClient() *redis.Client {
	return c.redis
}
