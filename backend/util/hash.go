package util

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func RandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func GenerateHash(password []byte, salt []byte) ([]byte, error) {
	return argon2.IDKey(password, salt, 3, 12288, 1, 32), nil
}
