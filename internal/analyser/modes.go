package analyser

import (
	"bufio"
	"log/slog"

	"github.com/symonk/log-analyse/internal/re"
)

type Strategy string

const (
	// keys for various strategy functions
	Sequential Strategy = "sequential"
	Reverse    Strategy = "reverse"
	FanOut     Strategy = "fanout"
)

var strategyMap = map[Strategy]func(loadedFile LoadedFile) []string{
	Sequential: sequentialFunc,
	Reverse:    reverseFunc,
	FanOut:     fanOutFunc,
}

// sequentialFunc processes a file sequentially
// using a single thread and reports matches against
// any of it's patterns
// TODO: This should put results on a channel, storing the
// entire file in memory is not going to cut it at scale!
func sequentialFunc(loadedFile LoadedFile) []string {
	defer loadedFile.File.Close()
	lines := make([]string, 0, 64)
	scanner := bufio.NewScanner(loadedFile.File)
	patterns, _ := re.CompileSlice(loadedFile.Options.Patterns)
	for scanner.Scan() {
		for _, pattern := range patterns {
			b := scanner.Bytes()
			if ok := pattern.Match(b); ok {
				line := string(scanner.Bytes())
				slog.Info("matched", slog.String("line", line), slog.Any("pattern", pattern))
				lines = append(lines, line)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		slog.Error("scanning error", slog.Any("error", err))
	}
	slog.Info("finished parsing file", slog.String("file", loadedFile.File.Name()))
	return lines
}

// reverseFunc traverses the file from the tail end backwards
// this is useful for finding later matches quicker.
func reverseFunc(loadedFile LoadedFile) []string {
	return nil
}

// fanOutFunc takes the current lines of the file, splits it
// into chunks and spawns multiple goroutines responsible for
// a smaller subset of the file and joins all matches back
// together.
func fanOutFunc(loadedFile LoadedFile) []string {
	return nil
}
