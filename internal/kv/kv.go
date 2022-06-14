package kv

import (
	"context"
)

// KV represents a key value store.
type KV interface {
	Close(ctx context.Context) error
}
