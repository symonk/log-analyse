package files

import (
	"github.com/symonk/log-analyse/internal/config"
)

type File struct {
}

type Flattener interface {
	Flatten([]config.FileConfig) ([]File, error)
}

type Option func(c *Condenser)

func WithStrict(strict bool) Option {
	return func(c *Condenser) {
		c.strict = true
	}
}

// Condenser takes an array of config location blocks
// and resolves them down to individual files.  It
// provides a flat set of files and config mapping that
// goes with each one of them.
type Condenser struct {
	locator Collector
	strict  bool
}

// NewCondenser returns a new instance (ptr) of
// Condenser
func NewCondenser(locator Collector, options ...Option) *Condenser {
	c := &Condenser{}
	for _, opt := range options {
		opt(c)
	}
	return c
}

// Flatten takes the config location blocks and flattens
// them into a configeration per found log file
func (c Condenser) Flatten([]config.FileConfig) ([]File, error) {

	return nil, nil
}
