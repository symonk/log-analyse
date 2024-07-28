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

func TestCanBuildValidConfig(t *testing.T) {
	cfg, valErr := loadAndValidateConfig(t, basic)
	assert.Nil(t, valErr)
	// TODO: implement proper checking, this is a sizable test!
	assert.Equal(t, cfg.Files[0].Options.Hits, 5)
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
	valErr := Validate(cfg)
	fmt.Println(valErr)
	expected := "Key: 'Config.Files' Error:Field validation for 'Files' failed on the 'required' tag"
	assert.ErrorContains(t, valErr, expected)
}

func TestGlobCannotBeEmpty(t *testing.T) {
	b := []byte(`
---
files:
  - glob: ""
  - glob: ""
  `)
	cfg, err := loadConfigFile(b)
	assert.Nil(t, err)
	valErr := Validate(cfg)
	expectedFirst := "Key: 'Config.Files[0].Glob' Error:Field validation for 'Glob' failed on the 'required' tag"
	expectedSecond := "Key: 'Config.Files[1].Glob' Error:Field validation for 'Glob' failed on the 'required' tag"
	fmt.Println(valErr)
	fmt.Println(valErr)
	assert.ErrorContains(t, valErr, expectedFirst)
	assert.ErrorContains(t, valErr, expectedSecond)
}

func TestOptionsAreRequired(t *testing.T) {
	b := []byte(`
---	
files:
  - glob: "foo.txt"

`)
	cfg, err := loadConfigFile(b)
	assert.Nil(t, err)
	valErr := Validate(cfg)
	expected := "Key: 'Config.Files[0].Options' Error:Field validation for 'Options' failed on the 'required' tag"
	assert.ErrorContains(t, valErr, expected)
}

func TestOptionsHitMustBeSuppliedAndPositiveInt(t *testing.T) {
	b := []byte(`
---
files:
  - glob: "foo.txt"
    options:
      hits: -1
`)
	_, valErr := loadAndValidateConfig(t, b)
	expectedNegative := "Key: 'Config.Files[0].Options.Hits' Error:Field validation for 'Hits' failed on the 'gt' tag"
	assert.ErrorContains(t, valErr, expectedNegative)
}

func TestPeriodMustBeValidTimeDuration(t *testing.T) {
	b := []byte(`
---
files:
  - glob: "foo.txt"
    options:
      hits: 1
      period: fail
`)
	_, valErr := loadAndValidateConfig(t, b)
	expected := "Key: 'Config.Files[0].Options.Period' Error:Field validation for 'Period' failed on the 'is-valid-time-duration' tag"
	assert.ErrorContains(t, valErr, expected)
}

func TestPatternsMustBeProvided(t *testing.T) {
	b := []byte(`
---
files:
  - glob: "foo.txt"
    options:
      hits: 1
      period: 10s
`)
	_, valErr := loadAndValidateConfig(t, b)
	expected := "Key: 'Config.Files[0].Options.Patterns' Error:Field validation for 'Patterns' failed on the 'required' tag"
	assert.ErrorContains(t, valErr, expected)

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

// loadAndValidateConfig unmarshals the byte stream through viper
// and into a config object, also applying validator validation
// against the input.
// Any validation errors are returned to the caller
func loadAndValidateConfig(t *testing.T, b []byte) (*Config, error) {
	cfg, err := loadConfigFile(b)
	if err != nil {
		t.Error(err)
	}
	valErr := Validate(cfg)
	return cfg, valErr

}
