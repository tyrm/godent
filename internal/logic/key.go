package logic

import "crypto/ed25519"

type Key interface {
	GenerateSigningKey() (ed25519.PrivateKey, error)
	GetPublicKey() (string, error)
}
