package monitor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/symonk/log-analyse/internal/files"
	"github.com/symonk/log-analyse/internal/re"
)

type Watcher interface {
	Watch(c files.ConfiguredFile, matches chan string)
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

// This is a hack right now. wip!
func (f *Filemon) Watch(c files.ConfiguredFile, done chan struct{}, matches chan string) {
	handle, err := os.Open(c.Path)
	if err != nil {
		// TODO: implement some solution
		panic(err)
	}
	patterns, _ := re.CompileSlice(c.Threshold.Patterns)
	strategyFn := strategyFactory[Matches]
	reader := bufio.NewReader(handle)
	fmt.Println("monitoring: ", c.Path)
	go func() {
		defer handle.Close()
		for {
			select {
			case <-done:
				return
			default:
				line, err := reader.ReadBytes('\n')
				if err != nil {
					if err == io.EOF {
						time.Sleep(time.Second)
						continue
					}
					// no bytes read maybe, various other errors
					fmt.Println(err)
					break
				}
				for _, p := range patterns {
					ok, match := strategyFn(line, p)
					if ok {
						matches <- fmt.Sprintf("match in file %s: %s", c.Path, match)
					}
				}

			}
		}
	}()

}
