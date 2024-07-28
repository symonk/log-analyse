package config

import (
	"github.com/go-playground/validator/v10"
)

const (
	configType = "yaml"
	configName = "log-analyse"
)

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
	return nil
}

// Config outlines the patterns of files to monitor
// and various around those files.
type Config struct {
	Files []FileConfig `yaml:"files" validate:"required,dive"`
}

func Validate(c *Config) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation(validDuration, isValidTimeDurationFunc)
	if err := validate.Struct(c); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err.(validator.ValidationErrors)

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
	Glob    string   `yaml:"glob" validate:"required"`
	Options *Options `yaml:"options" validate:"required"`
}

// Options encapsulates the configuration for each defined
// glob pattern in the config
type Options struct {
	Active   bool     `yaml:"active"`
	Hits     int      `yaml:"hits" validate:"gt=0"`
	Period   string   `yaml:"period" validate:"is-valid-time-duration"`
	Patterns []string `yaml:"patterns" validate:"required"`
	Notify   string   `yaml:"notify, omitempty"`
}
