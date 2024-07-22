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
func sequentialFunc(loadedFile LoadedFile) []string {
	defer loadedFile.File.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(loadedFile.File)
	patterns, _ := re.CompileSlice(loadedFile.Options.Patterns)
	for scanner.Scan() {
		line := scanner.Text()
		for _, pattern := range patterns {
			if ok := pattern.Match([]byte(line)); ok {
				slog.Info("matched", slog.String("line", line), slog.Any("pattern", pattern))
				lines = append(lines, line)
			}
		}
	}
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
