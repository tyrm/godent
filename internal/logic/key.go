package logic

import (
	"crypto/ed25519"
	"crypto/rand"
)

func (logic *Logic) GenerateSigningKey() (*ed25519.PrivateKey, error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	return &priv, err
}
