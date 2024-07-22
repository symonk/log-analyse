package analyser

import (
	"bufio"
	"log/slog"
	"regexp"
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
	for scanner.Scan() {
		line := scanner.Text()
		for _, pattern := range loadedFile.Threshold.Patterns {
			ok, err := regexp.Match(pattern, []byte(line))
			if err != nil {
				slog.Error("error matching line with pattern", slog.String("line", line), slog.String("pattern", pattern))
			}
			if ok {
				slog.Info("matched", slog.String("line", line), slog.String("pattern", pattern))
				lines = append(lines, line)
			}
		}
	}
	return lines
}
