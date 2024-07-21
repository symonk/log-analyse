package analyser

import (
	"os"

	"github.com/symonk/log-analyse/internal/config"
	"github.com/symonk/log-analyse/internal/files"
)

type Loader interface {
	Load() ([]LoadedFile, error)
}

type LoadedFile struct {
	File      *os.File
	Threshold config.Threshold
}

type FileLoader struct {
	files []files.IndividualFile
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
		loaded = append(loaded, LoadedFile{File: opened, Threshold: individualFile.Threshold})
	}
	return loaded, nil
}