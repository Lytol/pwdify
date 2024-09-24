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
	"errors"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

//go:embed template.html.tmpl
var templateContent string

var ValidExtensions = []string{".html"}

type Status struct {
	File  string
	Error error
}

func EncryptContent(path string, password string) ([]byte, error) {
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
	encrypted, err := EncryptContent(path, password)
	if err != nil {
		return err
	}

	content, err := protectedPageContent(encrypted)
	if err != nil {
		return err
	}

	return os.WriteFile(path, content, 0644)
}

func Encrypt(fileOrDir []string, password string) (chan Status, int, error) {
	ch := make(chan Status)

	files, err := expandFiles(fileOrDir)
	if err != nil {
		return ch, 0, err
	}

	// TODO: we are still serializing the file encryptions, which is a bit silly
	go func() {
		for _, file := range files {
			err := EncryptFile(file, password)
			ch <- Status{File: file, Error: err}
		}
		close(ch)
	}()

	return ch, len(files), err
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

func expandFiles(paths []string) ([]string, error) {
	files := []string{}
	errs := []error{}

	for _, path := range paths {
		fi, err := os.Stat(path)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if fi.IsDir() {
			expanded, err := getFiles(path)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			files = append(files, expanded...)
		} else {
			files = append(files, path)
		}
	}

	return files, errors.Join(errs...)
}

func getFiles(dir string) ([]string, error) {
	files := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if isValidExtension(path) {
			files = append(files, path)
		}
		return err
	})

	return files, err
}

func isValidExtension(path string) bool {
	for _, ext := range ValidExtensions {
		if filepath.Ext(path) == ext {
			return true
		}
	}
	return false
}
