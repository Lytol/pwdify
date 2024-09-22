package pwdify

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"
)

func Encrypt(path string, password string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Generate a key from the password using SHA-256
	key := sha256.Sum256([]byte(password))

	// Encrypt the data using AES
	return encryptData(data, key[:])
}

func EncryptFile(path string, password string) error {
	encrypted, err := Encrypt(path, password)
	if err != nil {
		return err
	}

	// TODO: Generate from template

	return os.WriteFile(path, encrypted, 0644)
}

func encryptData(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}
