package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GlobalConfig *Config

func Get() *Config {
	return GlobalConfig
}

// Init Config loads the config into memory
func Init(configFilePath string) {
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
	} else {
		baseDir, err := defaultConfigPath()
		cobra.CheckErr(err)

		viper.AddConfigPath(baseDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("log-analyse")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err := viper.Unmarshal(&GlobalConfig); err != nil {
			slog.Error("configuration file was not valid", slog.String("config", viper.ConfigFileUsed()), slog.Any("error", err))
			os.Exit(2)
		}
	} else {
		slog.Error("no config file could be found: ", slog.Any("error", err))
		os.Exit(1)
	}
	slog.Info("Successfully built a config")
}

// Config encapsulates basic configuration used by
// log-analyse when it runs.  At the moment this is
// relatively basic, enhanced configurations will be
// enabled in future.
type Config struct {
	Files        []FileConfig  `yaml:"files"`
	Integrations []Integration `yaml:"integrations, omitempty"`
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
	Glob        string      `yaml:"glob" validate:"required"`
	Options     Options     `yaml:"Options" validate:"required"`
	Integration Integration `yaml:"Integration, omitempty"`
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

// Integration is an implementation of an alerting
// mechanism
type Integration struct {
	Slack Slack `yaml:"slack, omitempty"`
	Email Email `yaml:"email, omitempty"`
}

// Slack encapsulates configurations for the slack
// notification plugin
type Slack struct {
	Webhook string `yaml:"webhook"`
}

// Email encapsulates configurations for the email
// notification plugin
type Email struct {
	To []string `yaml:"to"`
}
