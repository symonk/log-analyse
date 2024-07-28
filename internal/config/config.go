package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
	configName = "log-analyse"
)

var GlobalConfig *Config

func Get() *Config {
	return GlobalConfig
}

// Init Config loads the config into memory
func Init(configFilePath string) error {
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
	} else {
		baseDir, err := defaultConfigPath()
		cobra.CheckErr(err)

		viper.AddConfigPath(baseDir)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
	}

	if err := viper.ReadInConfig(); err == nil {
		if err := viper.Unmarshal(&GlobalConfig); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

// Config outlines the patterns of files to monitor
// and various around those files.
type Config struct {
	Files []FileConfig `yaml:"files"`
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
	Glob    string  `yaml:"glob" validate:"required"`
	Options Options `yaml:"Options" validate:"required"`
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
