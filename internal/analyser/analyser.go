package analyser

import (
	"github.com/symonk/log-analyse/internal/config"
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
	fileConfigs []config.FileConfig
	strategy    string
	maxBound    int
}

// NewFileAnalyser returns a new file analyser.
func NewFileAnalyser(fileConfigs []config.FileConfig, options ...Option) *FileAnalyser {
	analyser := &FileAnalyser{fileConfigs: fileConfigs}
	for _, option := range options {
		option(analyser)
	}
	return analyser
}
