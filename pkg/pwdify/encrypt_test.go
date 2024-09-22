package pwdify_test

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lytol/pwdify/pkg/pwdify"
)

func TestEncrypt(t *testing.T) {
	encrypted, err := pwdify.Encrypt("testdata/example.html", "qwerty123")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", hex.EncodeToString(encrypted))
}

func TestEncryptFile(t *testing.T) {
	// tmpDir := t.TempDir()
	// tmpPath := filepath.Join(tmpDir, "example.html")
	tmpPath := filepath.Join("testdata", "_example.html")

	err := copy("testdata/example.html", tmpPath)
	if err != nil {
		t.Fatal(err)
	}

	if err := pwdify.EncryptFile(tmpPath, "qwerty123"); err != nil {
		t.Fatal(err)
	}
}

func copy(src, dest string) error {
	contents, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dest, contents, 0644)
}
