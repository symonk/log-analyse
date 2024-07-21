package analyser

import (
	"bufio"
	"log/slog"
	"regexp"
	"sync"

	"github.com/symonk/log-analyse/internal/files"
)

// Result is a matched result
type Result string

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
	Analyse() (chan<- Result, error)
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
// TODO: optimise patterns i.e reuse compiled
func (f *FileAnalyser) Analyse() (<-chan string, error) {
	loadedFiles, err := f.loader.Load()
	if err != nil {
		// TODO: no good!
		panic(err)
	}

	results := make(chan string)
	work := make(chan Task)
	_ = work

	var wg sync.WaitGroup
	wg.Add(len(loadedFiles))

	// TODO: spawn the workers etc.

	// shove the loaded files onto the channel for processing

	for _, f := range loadedFiles {
		go func() {
			defer wg.Done()
			defer f.File.Close()
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
						results <- line
					}
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	return results, nil
}
