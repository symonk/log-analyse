package files

type Collector interface {
	Locate() error
}

type FileCollector struct {
}

func NewFileLocator() *FileCollector {
	return &FileCollector{}
}

func (f FileCollector) Locate() error {
	return nil
}
