package analyser

var strategyMap = map[string]func(loadedFile LoadedFile) []string{
	"sequential": seqScanStrategyFn,
}

// seqScanStrategyFn processes a file sequentially
// using a single thread and reports matches against
// any of it's patterns
func seqScanStrategyFn(loadedFile LoadedFile) []string {
	return nil
}
