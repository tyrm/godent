package models

import (
	"crypto/aes"
	gocipher "crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"strings"

)

var errDataTooSmall = errors.New("data too small")

func decrypt(b []byte) ([]byte, error) {
	l := logger.WithField("func", "decrypt")

	gcm, err := getCrypto()
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(b) < nonceSize {
		l.Error(errDataTooSmall.Error())

		return nil, errDataTooSmall
	}

	nonce, ciphertext := b[:nonceSize], b[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		l.Errorf("decrypting: %s", err.Error())

		return nil, err
	}

	return plaintext, nil
}

func encrypt(b []byte) ([]byte, error) {
	l := logger.WithField("func", "encrypt")

	gcm, err := getCrypto()
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		l.Errorf("reading nonce: %s", err.Error())

		return nil, err
	}

	return gcm.Seal(nonce, nonce, b, nil), nil
}

func getCrypto() (gocipher.AEAD, error) {
	l := logger.WithField("func", "getCrypto").WithField("type", "EncryptedString")

	key := []byte(strings.ToLower(viper.GetString(config.Keys.DBEncryptionKey)))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		l.Errorf("new cipher: %s", err.Error())

		return nil, err
	}

	gcm, err := gocipher.NewGCM(cipher)
	if err != nil {
		l.Errorf("new gcm: %s", err.Error())

		return nil, err
	}

	return gcm, nil
}
