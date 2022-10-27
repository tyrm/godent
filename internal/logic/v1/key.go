package v1

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"github.com/tyrm/godent/internal/logic"
)

func (logic *Logic) GenerateSigningKey() (ed25519.PrivateKey, error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)

	return priv, err
}

func (logic *Logic) GetPublicKey() (string, error) {
	pubKey, ok := logic.signingKey.Public().(ed25519.PublicKey)
	if !ok {
		return "", errors.New("public key not ed25519")
	}

	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(pubKey), nil
}

func (logic *Logic) IsEphemeralPubKeyValid(ctx context.Context, pubKey string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (logic *Logic) IsPubKeyValid(ctx context.Context, pubKey string) (bool, error) {
	publicKey, err := logic.GetPublicKey()
	if err != nil {
		return false, err
	}

	return pubKey == publicKey, nil
}

func getPrivateKey() (ed25519.PrivateKey, error) {
	l := logger.WithField("func", "getPrivateKey")

	signingKey := viper.GetString(config.Keys.SigningKey)
	if signingKey == "" {
		return nil, logic.ErrNotFound
	}

	keyParts := strings.Split(signingKey, " ")
	if len(keyParts) != 3 {
		return nil, logic.ErrInvalid
	}

	l.Tracef("parsing key: %s", keyParts[2])

	privateKey, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(keyParts[2])
	if err != nil {
		return nil, err
	}

	return ed25519.NewKeyFromSeed(privateKey), nil
}
