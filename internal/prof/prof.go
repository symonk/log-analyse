package prof

import (
	"os"
	"runtime/pprof"
)

func RunProf() func() {
	f, err := os.Create("analyse.prof")
	if err != nil {
		os.Exit(1)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		os.Exit(1)
	}
	return pprof.StopCPUProfile
}
