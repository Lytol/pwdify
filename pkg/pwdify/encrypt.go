package pwdify

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"html/template"
	"io"
	"os"
)

//go:embed template.html.tmpl
var templateContent string

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

	content, err := protectedPageContent(encrypted)
	if err != nil {
		return err
	}

	return os.WriteFile(path, content, 0644)
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

	full := gcm.Seal(nonce, nonce, data, nil)

	return full, nil
}

func protectedPageContent(data []byte) ([]byte, error) {
	var content bytes.Buffer
	w := bufio.NewWriter(&content)

	tmpl, err := template.New("protectedPage").Parse(templateContent)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(w, struct{ EncryptedContent string }{hex.EncodeToString(data)})
	if err != nil {
		return nil, err
	}

	err = w.Flush()
	if err != nil {
		return nil, err
	}

	return content.Bytes(), nil
}
