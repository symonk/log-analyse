package monitor

type Watcher interface{}

// Filemon is responsible for tailing an arbitrary number of files
type Filemon struct {
}

func (f *Filemon) Watch(path string) {

}
