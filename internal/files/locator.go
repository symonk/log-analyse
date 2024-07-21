package files

import (
	"log/slog"
	"path/filepath"

	"github.com/symonk/log-analyse/internal/config"
)

type IndividualFile struct {
	Path      string
	Threshold config.Threshold
}

type Collector interface {
	Locate() ([]IndividualFile, error)
}

// FileCollector is responsible for taking globs and their
// paired configuration and flattening those glob patterns
// into actual actionable files and the configs that apply
// to them
type FileCollector struct {
	cfg *config.Config
}

func NewFileLocator(cfg *config.Config) *FileCollector {
	return &FileCollector{cfg: cfg}
}

func (f FileCollector) Locate() ([]IndividualFile, error) {
	files := make([]IndividualFile, 0)
	for _, file := range f.cfg.Files {
		flattened, err := f.filesFromGlob(file.Glob)
		if err != nil {
			return files, err
		}
		if len(flattened) == 0 {
			slog.Warn("no files for glob", "glob", file.Glob)
			return files, nil
		}
		for _, f := range flattened {
			files = append(files, IndividualFile{Path: f, Threshold: file.Threshold})
		}
	}
	// TODO: Handle duplicate paths here; multiple config blocks can overlap
	// file matches, not sure if we care (yet) as each file will have it's
	// proper configuration associated with it, we may just read the same file
	// n times applying a different config, that might be ok!
	return files, nil
}

// filesFromGlob takes a glob pattern and returns the slice of file paths
// matching that glob.
func (f FileCollector) filesFromGlob(glob string) ([]string, error) {
	return filepath.Glob(glob)
}
