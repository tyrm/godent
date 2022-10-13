package v1

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"strings"

	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/logic"
)

func (logic *Logic) GenerateSigningKey() (ed25519.PrivateKey, error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)

	return priv, err
}

func getPrivateKey() (ed25519.PrivateKey, error) {
	signingKey := viper.GetString(config.Keys.SigningKey)
	if signingKey == "" {
		return nil, logic.ErrNotFound
	}

	keyParts := strings.Split(signingKey, " ")
	if len(keyParts) != 3 {
		return nil, logic.ErrInvalid
	}

	privateKey, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(keyParts[2])
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
