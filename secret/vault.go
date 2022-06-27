package secret

import (
	"errors"

	"github.com/sanketrai1/vault/encrypt"
)

type store interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
}

type vault struct {
	key     string
	storage map[string]string
}

func Vault(key string) (vault, error) {
	v := vault{key: key, storage: make(map[string]string)}
	return v, nil
}

func (v *vault) Set(key string, value string) (string, error) {
	cipher, err := encrypt.Encrypt(v.key, value)
	if err != nil {
		return "", err
	}
	v.storage[key] = cipher
	return cipher, nil
}

func (v *vault) Get(key string) (string, error) {
	if val, ok := v.storage[key]; ok {
		plaintext, err := encrypt.Decrypt(v.key, val)
		if err != nil {
			return "", err
		}
		return plaintext, nil
	}
	return "", errors.New("Key not found")
}
