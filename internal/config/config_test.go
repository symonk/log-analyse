package config

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var basic = []byte(`
---
files:
  - glob: "~/logs/*.txt"
    options:
      hits: 5
      period: 30s
      patterns:
        - ".*FATAL.*"
        - ".*payment failed.*"
      notify: "email"

  - glob: "~/logs/foo.log"
    options:
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
    options:
      hits: 5
      period: 30s
      patterns:
        - ".*FATAL.*"
        - ".*payment failed.*"
      notify: "email"
`)

var empty = []byte(`
---
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

func TestCanLoadConfigFromAbsolutePath(t *testing.T) {

}

func TestTopLevelFilesIsRequired(t *testing.T) {
	b := []byte(``)
	cfg, err := loadConfigFile(b)
	assert.Nil(t, err)
	valErr := cfg.Validate()
	fmt.Println(valErr)
	expected := `Key: 'Config.Files' Error:Field validation for 'Files' failed on the 'required' tag`
	assert.ErrorContains(t, valErr, expected)
}

func TestFileConfigMustHaveAGlob(t *testing.T) {
	var b = []byte(`
---
files:
  - glob: ""
  `)
	cfg, err := loadConfigFile(b)
	assert.Nil(t, err)
	valErr := cfg.Validate()
	assert.ErrorIs(t, valErr, nil)

}

func TestOptionsAreRequired(t *testing.T) {
	var b = []byte(`
---
files:
  - glob: "ok.txt"
    options: 
  `)
	cfg, err := loadConfigFile(b)
	assert.Nil(t, err)
	valErr := cfg.Validate()
	assert.ErrorIs(t, valErr, nil)

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
