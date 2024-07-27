package monitor

import "regexp"

type Strategy string

type StrategyFunc func(bytes []byte, regexp *regexp.Regexp) (bool, string)

const (
	Matches Strategy = "matches"
)

var strategyFactory = map[Strategy]StrategyFunc{
	Matches: matches,
}

func matches(bytes []byte, regexp *regexp.Regexp) (bool, string) {
	if regexp.Match(bytes) {
		return true, string(bytes)
	}
	return false, ""
}
