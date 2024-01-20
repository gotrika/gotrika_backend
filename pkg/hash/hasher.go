package hash

import (
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

type Hasher interface {
	Hash(value string) (string, error)
}

type ScryptHasher struct {
	secretKey string
}

func NewScryptHasher(secretKey string) *ScryptHasher {
	return &ScryptHasher{secretKey: secretKey}
}

func (h *ScryptHasher) Hash(value string) (string, error) {
	salt := []byte(h.secretKey)
	dk, err := scrypt.Key([]byte(value), salt, 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	hash := base64.StdEncoding.EncodeToString(dk)
	return hash, nil
}
