package analyser

import (
	"bufio"
	"log/slog"
	"regexp"

	"github.com/symonk/log-analyse/internal/files"
)

// Option is a functional option for the Analyser
type Option func(a *FileAnalyser)

// WithStrategy allows an arbitrary strategy to be
// applied to all file analysing.
func WithStrategy(strategy string) Option {
	return func(a *FileAnalyser) {
		a.strategy = strategy
	}
}

// WithBounds applies an upper limit to the total number
// of goroutines the analyser is allowed to spawn for
// completing its work.  Default is 0 (no limit)
func WithBounds(bound int) Option {
	return func(a *FileAnalyser) {
		a.maxBound = bound
	}
}

// Analyser is the interface for something that can analyse
// patterns against files on disk
type Analyser interface {
	Analyse() error
}

// FileAnalyser accepts file paths with their paired thresholds
// and is responsible for distributing the workloads for
// files based on the user configuration and funneling errors
// back to the calling commands
// The files are already flattened by the time the analyser is
// called to do its work.
// Analysrer is not a live tailer, but a retrospective log inspector
// of sorts.
type FileAnalyser struct {
	flattenedFiles []files.IndividualFile
	strategy       string
	maxBound       int
	loader         *FileLoader
}

// NewFileAnalyser returns a new file analyser.
func NewFileAnalyser(individualFiles []files.IndividualFile, options ...Option) *FileAnalyser {
	analyser := &FileAnalyser{flattenedFiles: individualFiles, loader: &FileLoader{files: individualFiles}}
	for _, option := range options {
		option(analyser)
	}
	return analyser
}

// Analyse performs retrospect log file analysis
// TODO: doing more than 1 thing
// TODO: abstract file loading
// TODO: abstract file scanning
// TODO: optimise patterns i.e reuse compiled
func (f *FileAnalyser) Analyse() error {
	// TODO: check files exist, do what we can or add a strict flag

	// TODO: Asynchronously process all files, all lines scaling out massively
	// TODO: matching patterns in the config file
	loadedFiles, err := f.loader.Load()
	if err != nil {
		// TODO: no good!
		panic(err)
	}
	for _, f := range loadedFiles {
		scanner := bufio.NewScanner(f.File)
		for scanner.Scan() {
			line := scanner.Text()
			for _, pattern := range f.Threshold.Patterns {
				ok, err := regexp.Match(pattern, []byte(line))
				if err != nil {
					slog.Error("error matching line with pattern", slog.String("line", line), slog.String("pattern", pattern))
				}
				if ok {
					slog.Info("matched", slog.String("line", line), slog.String("pattern", pattern))
				}
			}
		}
	}
	return nil
}
