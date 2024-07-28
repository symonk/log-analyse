package files

import (
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/symonk/log-analyse/internal/config"
)

type NoFilesFromGlobError struct {
	glob string
}

func (n *NoFilesFromGlobError) Error() string {
	return fmt.Sprintf("pattern %q did not have any files matched", n.glob)
}

type ConfiguredFile struct {
	Path      string
	Threshold config.Options
}

type Collector interface {
	Locate() ([]ConfiguredFile, error)
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

func (f FileCollector) Locate() ([]ConfiguredFile, error) {
	monitorableFiles := make([]ConfiguredFile, 0)
	for _, file := range f.cfg.Files {
		if !file.Options.Active {
			continue
		}
		flattened, err := f.filesFromGlob(file.Glob)
		if err != nil {
			return monitorableFiles, err
		}
		if len(flattened) == 0 {
			slog.Warn("no files for glob", "glob", file.Glob)
			return monitorableFiles, &NoFilesFromGlobError{glob: file.Glob}
		}
		for _, f := range flattened {
			monitorableFiles = append(monitorableFiles, ConfiguredFile{Path: f, Threshold: *file.Options})
		}
	}
	// TODO: Handle duplicate paths here; multiple config blocks can overlap
	// file matches, not sure if we care (yet) as each file will have it's
	// proper configuration associated with it, we may just read the same file
	// n times applying a different config, that might be ok!
	return monitorableFiles, nil
}

// filesFromGlob takes a glob pattern and returns the slice of file paths
// matching that glob.
func (f FileCollector) filesFromGlob(glob string) ([]string, error) {
	return filepath.Glob(glob)
}
