package config

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var basic = []byte(`
---
files:
  # A glob based folder lookup
  - glob: "~/logs/*.txt"
    threshold:
      hits: 5
      period: 30s
      patterns:
        - ".*FATAL.*"
        - ".*payment failed.*"
      notify: "email"
  # An explicit log file
  - glob: "~/logs/foo.log"
    threshold:
      hits: 1
      period: 1m
    patterns:
      - ".*disk space low.*"
    notify: "slack"

`)

func TestCanUnmarshalConfigSuccessfully(t *testing.T) {
	c := loadConfigFile(t, basic)
	assert.Len(t, c.Files, 2)
}

func TestReturnGlobs(t *testing.T) {
	fConfigs := []FileConfig{{Glob: "foo"}, {Glob: "bar"}, {Glob: "baz"}}
	c := Config{Files: fConfigs}
	assert.Equal(t, c.Globs(), []string{"foo", "bar", "baz"})
}

// loadConfigFile streams a byte slice into a viper config and
// unmarshals it into the Config object.
func loadConfigFile(t *testing.T, b []byte) *Config {
	var c *Config
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(b))
	if err := viper.Unmarshal(&c); err != nil {
		t.Fatal("could not load config")
	}
	return c
}
