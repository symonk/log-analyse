package files

type Collector interface {
	Locate() error
}

// FileCollector is responsible for taking globs and their
// paired configuration and flattening those glob patterns
// into actual actionable files and the configs that apply
// to them
type FileCollector struct {
}

func NewFileLocator() *FileCollector {
	return &FileCollector{}
}

func (f FileCollector) Locate() error {
	return nil
}
