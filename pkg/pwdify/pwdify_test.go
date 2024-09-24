package pwdify_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lytol/pwdify/pkg/pwdify"
)

func TestEncryptFile(t *testing.T) {
	tmpPath := filepath.Join("testdata", "_example.html")

	err := copy("testdata/example.html", tmpPath)
	if err != nil {
		t.Fatal(err)
	}

	if err := pwdify.EncryptFile(tmpPath, "qwerty123"); err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Fatal(err)
	}

	// TODO: actually assert SOMETHING
}

func copy(src, dest string) error {
	contents, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dest, contents, 0644)
}
