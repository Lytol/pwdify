package pwdify

import "time"

type Status struct {
	File  string
	Error error
}

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Run(files []string, password string) chan Status {
	ch := make(chan Status)

	go func() {
		for _, file := range files {
			time.Sleep(1 * time.Second)
			err := EncryptFile(file, password)
			ch <- Status{File: file, Error: err}
		}
		close(ch)
	}()

	return ch
}
