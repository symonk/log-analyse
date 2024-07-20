package files

type Locator interface {
	Locate() error
}

type FileLocator struct {
}

func NewFileLocator() *FileLocator {
	return &FileLocator{}
}

func (f FileLocator) Locate() error {
	return nil
}
