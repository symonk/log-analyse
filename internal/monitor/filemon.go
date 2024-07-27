package monitor

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
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

// This is a hack right now. wip!
func (f *Filemon) Watch(c files.ConfiguredFile, done chan struct{}) {
	handle, err := os.Open(c.Path)
	if err != nil {
		// TODO: implement some solution
		panic(err)
	}
	defer handle.Close()
	patterns, _ := re.CompileSlice(c.Threshold.Patterns)
	strategyFn := strategyFactory[Matches]
	_, _ = patterns, strategyFn
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watch.Close()
	if err := watch.Add("/tmp"); err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case <-done:
				return
			case event, ok := <-watch.Events:
				if !ok {
					return
				}
				fmt.Println(event, ok)
			case err, ok := <-watch.Errors:
				if !ok {
					return
				}
				fmt.Println(err, ok)
			}
		}

	}()
}
