package analyser

import (
	"bufio"
	"log/slog"

	"github.com/symonk/log-analyse/internal/re"
)

var strategyMap = map[string]func(loadedFile LoadedFile) []string{
	"sequential": seqScanStrategyFn,
}

// seqScanStrategyFn processes a file sequentially
// using a single thread and reports matches against
// any of it's patterns
func seqScanStrategyFn(loadedFile LoadedFile) []string {
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
