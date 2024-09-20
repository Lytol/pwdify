package util

import (
	"fmt"
	"os"
)

const DefaultLogFilename = "debug.log"

type Logger interface {
	Logf(string, ...interface{})
	Close()
}

type DiscardLogger struct{}

func (DiscardLogger) Logf(string, ...interface{}) {
	// Do nothing
}

func (DiscardLogger) Close() {
	// Do nothing
}

type FileLogger struct {
	file *os.File
}

func (l FileLogger) Logf(format string, args ...interface{}) {
	fmt.Fprintf(l.file, format, args...)
}

func (l FileLogger) Close() {
	l.file.Close()
}

func NewLogger() (Logger, error) {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := os.OpenFile(DefaultLogFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		return FileLogger{f}, err
	} else {
		return DiscardLogger{}, nil
	}
}
