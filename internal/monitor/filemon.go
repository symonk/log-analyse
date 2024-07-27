package monitor

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/symonk/log-analyse/internal/files"
	"github.com/symonk/log-analyse/internal/re"
)

type Watcher interface {
	Watch(c files.ConfiguredFile)
}

type Opener interface {
	Open(path string) (*os.File, error)
}

// Filemon is responsible for tailing an arbitrary number of files
type Filemon struct {
	opener Opener
}

func NewFilemon() *Filemon {
	return &Filemon{}
}

func (f *Filemon) Watch(c files.ConfiguredFile) {
	handle, err := os.Open(c.Path)
	if err != nil {
		// TODO: implement some solution
		panic(err)
	}
	defer handle.Close()
	reader := bufio.NewReader(handle)
	patterns, _ := re.CompileSlice(c.Threshold.Patterns)
	strategyFn := strategyFactory[Matches]
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				time.Sleep(time.Second)
				continue
			}
			break
		}
		for _, p := range patterns {
			ok, match := strategyFn(line, p)
			if ok {
				fmt.Print(match)
			}
		}
	}
}
