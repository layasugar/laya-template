package rotatelog

import (
	"errors"
	"os"
	"path/filepath"
)

type BuffIO struct {
	file *os.File
}

// GenerateBuffIO CreateFile creates a new fh in the given path, creating parent directories
// as necessary
func GenerateBuffIO() IOWriter {
	return new(BuffIO)
}

// CreateFile creates a new fh in the given path, creating parent directories
// as necessary
func (d *BuffIO) CreateFile(filename string) error {
	// make sure the dir is existed, eg:
	// ./foo/bar/baz/hello.logger must make sure ./foo/bar/baz is existed
	dirname := filepath.Dir(filename)
	if err := os.MkdirAll(dirname, 0755); err != nil {
		return err
	}
	// if we got here, then we need to create a fh
	fh, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	d.file = fh
	return nil
}

func (d *BuffIO) SyncFile() error {
	if d.file == nil {
		return errors.New("file no open")
	}

	return d.file.Sync()
}

func (d *BuffIO) Write(data []byte) (int, error) {
	if d.file == nil {
		return 0, nil
	}
	return d.file.Write(data)
}

func (d *BuffIO) Close() error {
	if d == nil {
		return nil
	}
	return d.file.Close()
}
