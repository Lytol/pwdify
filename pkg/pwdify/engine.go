package pwdify

import (
	"errors"
	"os"
	"path/filepath"
)

var ValidExtensions = []string{".html"}

type Status struct {
	File  string
	Error error
}

type Engine struct {
	Password string
	Files    []string
}

func New(fileOrDir []string, password string) (*Engine, error) {
	var errs []error

	e := &Engine{
		Password: password,
	}

	for _, path := range fileOrDir {
		fi, err := os.Stat(path)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if fi.IsDir() {
			files, err := getFiles(path)
			if err != nil {
				return nil, err
			}
			e.Files = append(e.Files, files...)
		} else {
			e.Files = append(e.Files, path)
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return e, nil
}

func (e *Engine) Run() chan Status {
	ch := make(chan Status)

	// TODO: we are still serializing the file encryptions, which is a bit silly
	go func() {
		for _, file := range e.Files {
			err := EncryptFile(file, e.Password)
			ch <- Status{File: file, Error: err}
		}
		close(ch)
	}()

	return ch
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
