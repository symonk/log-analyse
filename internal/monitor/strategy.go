package monitor

type Strategy string

type StrategyFunc func(bytes []byte) (bool, string)

const (
	Matches Strategy = "matches"
)

var strategyFactory = map[Strategy]StrategyFunc{
	Matches: matches,
}

func matches(bytes []byte) (bool, string) {

	return true, ""
}
