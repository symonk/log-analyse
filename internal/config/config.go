package config

import (
	"fmt"
)

const (
	configType = "yaml"
	configName = "log-analyse"
)

// Validatable is implemented by config types that
// guarantee required config is in place at runtime
// as well as handle default cases.
type Validatable interface {
	Validate() error
}

var GlobalConfig *Config

func Get() *Config {
	return GlobalConfig
}

// Init unmarshals the user provided config data or if omitted
// looks up the config in the default location.  Init also handles
// unmarshalling the config through viper into the config object.
// Additional config validation is finally performed.
func Init(configFilePath string) error {
	if err := parseViper(configFilePath); err != nil {
		return err
	}
	if err := GlobalConfig.Validate(); err != nil {
		return err
	}
	return nil
}

// Config outlines the patterns of files to monitor
// and various around those files.
type Config struct {
	Files []FileConfig `yaml:"files" validate: "required"`
}

// Validates ensures the unmarshalled config is fit
// for purposes and conforms to an appropriate standard
// for execution
func (c *Config) Validate() error {
	if len(c.Files) == 0 {
		return ErrFilesRequired
	}
	for _, fc := range c.Files {
		return fc.Validate()
	}
	return nil
}

// Globs returns the configured glob patterns defined in
// the configuration file.
func (c Config) Globs() []string {
	globs := make([]string, 0, len(c.Files))
	for _, file := range c.Files {
		globs = append(globs, file.Glob)
	}
	return globs
}

// FileConfig encapsualates the threshold for pattern
// matches before an alert or action is triggered.
type FileConfig struct {
	Glob    string   `yaml:"glob"`
	Options *Options `yaml:"Options"`
}

func (f *FileConfig) Validate() error {
	if f.Glob == "" {
		return ErrGlobRequired
	}
	if f.Options == nil {
		return ErrOptionsRequired
	}
	return nil
}

// Options encapsulates the configuration for each defined
// glob pattern in the config
type Options struct {
	Active   bool     `yaml:"active"`
	Hits     int      `yaml:"hits"`
	Period   string   `yaml:"period"`
	Patterns []string `yaml:"patterns"`
	Notify   string   `yaml:"notify, omitempty"`
}

func (o *Options) Validate() error {
	if o.Hits <= 0 {
		return fmt.Errorf("`Option.Hits` must be specified and be greater than 0")
	}
	return nil
}
