package files

import (
	"os"

	"github.com/symonk/log-analyse/internal/config"
)

type Loader interface {
	Load() ([]LoadedFile, error)
}

type LoadedFile struct {
	File    *os.File
	Options config.Options
}

type FileLoader struct {
	files []IndividualFile
}

// Load opens file paths and returns pointers to real files
// TODO: Error handling needs some improvements.
func (f *FileLoader) Load() ([]LoadedFile, error) {
	loaded := make([]LoadedFile, 0, len(f.files))
	for _, individualFile := range f.files {
		opened, err := os.Open(individualFile.Path)
		if err != nil {
			return loaded, err
		}
		loaded = append(loaded, LoadedFile{File: opened, Options: individualFile.Threshold})
	}
	return loaded, nil
}
