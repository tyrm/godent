package logic

import (
	"context"
	"crypto/ed25519"
)

type Key interface {
	GenerateSigningKey() (ed25519.PrivateKey, error)
	GetPublicKey() (string, error)
	IsEphemeralPubKeyValid(ctx context.Context, pubKey string) (bool, error)
	IsPubKeyValid(ctx context.Context, pubKey string) (bool, error)
}
