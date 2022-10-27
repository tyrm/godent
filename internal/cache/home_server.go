package cache

import "context"

type HomeServer interface {
	DeleteHomeServer(ctx context.Context, domain string) error
	GetHomeServer(ctx context.Context, domain string) (string, error)
	SetHomeServer(ctx context.Context, domain, homeServer string, expiration int) error
}
