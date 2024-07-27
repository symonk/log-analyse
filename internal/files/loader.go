package files

import (
	"os"

	"github.com/symonk/log-analyse/internal/config"
)

type Loader interface {
	Load() ([]OpenedFile, error)
}

type OpenedFile struct {
	File    *os.File
	Options config.Options
}

// Close closes the underlying file handle
// and returns any error.
func (o *OpenedFile) Close() error {
	return o.File.Close()
}

type FileLoader struct {
	files []ConfiguredFile
}

// Load opens file paths and returns pointers to real files
// TODO: Error handling needs some improvements.
func (f *FileLoader) Load() ([]OpenedFile, error) {
	loaded := make([]OpenedFile, 0, len(f.files))
	for _, individualFile := range f.files {
		opened, err := os.Open(individualFile.Path)
		if err != nil {
			return loaded, err
		}
		loaded = append(loaded, OpenedFile{File: opened, Options: individualFile.Threshold})
	}
	return loaded, nil
}
