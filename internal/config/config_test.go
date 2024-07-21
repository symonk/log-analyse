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
  - glob: "~/logs/*.txt"
    threshold:
      hits: 5
      period: 30s
      patterns:
        - ".*FATAL.*"
        - ".*payment failed.*"
      notify: "email"

  - glob: "~/logs/foo.log"
    threshold:
      hits: 1
      period: 1m
    patterns:
      - ".*disk space low.*"
    notify: "slack"

`)

var single = []byte(`
---
files:
  - glob: "~/logs/*.txt"
    threshold:
      hits: 5
      period: 30s
      patterns:
        - ".*FATAL.*"
        - ".*payment failed.*"
      notify: "email"
`)

func TestCanUnmarshalConfigSuccessfully(t *testing.T) {
	c, err := loadConfigFile(basic)
	assert.Nil(t, err)
	assert.Len(t, c.Files, 2)
}

func TestCanLoadSingleConfigBlock(t *testing.T) {
	c, err := loadConfigFile(single)
	assert.Nil(t, err)
	assert.Len(t, c.Files, 1)
}

func TestReturnGlobs(t *testing.T) {
	fConfigs := []FileConfig{{Glob: "foo"}, {Glob: "bar"}, {Glob: "baz"}}
	c := Config{Files: fConfigs}
	assert.Equal(t, c.Globs(), []string{"foo", "bar", "baz"})
}

// loadConfigFile streams a byte slice into a viper config and
// unmarshals it into the Config object.
func loadConfigFile(b []byte) (*Config, error) {
	var c *Config
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
		return c, err
	}
	if err := viper.Unmarshal(&c); err != nil {
		return c, err
	}
	return c, nil
}
