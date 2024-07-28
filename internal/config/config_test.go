package config

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var full = []byte(`
---
files:
  - glob: ~/logs/*.log
    options:
      active: false
      hits: 5
      period: 30s
      trigger: email
      patterns:
        - .*FATAL.*
        - .*payment failed.*
  - glob: ~/logs/foo.log
    options:
      active: true
      hits: 1
      period: 1h10s
      trigger: slack
      patterns:
        - .*critical error.*
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

func TestCanBuildValidConfig(t *testing.T) {
	cfg, err := loadAndValidateConfig(t, full)
	assert.Nil(t, err)
	assert.Equal(t, cfg.Files[0].Glob, "~/logs/*.log")
	assert.Equal(t, cfg.Files[1].Glob, "~/logs/foo.log")
	assert.Equal(t, cfg.Files[0].Options.Active, false)
	assert.Equal(t, cfg.Files[1].Options.Active, true)
	assert.Equal(t, cfg.Files[0].Options.Hits, 5)
	assert.Equal(t, cfg.Files[1].Options.Hits, 1)
	assert.Equal(t, cfg.Files[0].Options.Period, "30s")
	assert.Equal(t, cfg.Files[1].Options.Period, "1h10s")
	assert.Equal(t, cfg.Files[0].Options.Patterns, []string{".*FATAL.*", ".*payment failed.*"})
	assert.Equal(t, cfg.Files[1].Options.Patterns, []string{".*critical error.*"})
	assert.Equal(t, cfg.Files[0].Options.Trigger, "email")
	assert.Equal(t, cfg.Files[1].Options.Trigger, "slack")
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
	_, err := loadAndValidateConfig(t, b)
	expected := "Key: 'Config.Files' Error:Field validation for 'Files' failed on the 'required' tag"
	assert.ErrorContains(t, err, expected)
}

func TestGlobCannotBeEmpty(t *testing.T) {
	b := []byte(`
---
files:
  - glob: ""
  - glob: ""
  `)
	_, err := loadAndValidateConfig(t, b)
	expectedFirst := "Key: 'Config.Files[0].Glob' Error:Field validation for 'Glob' failed on the 'required' tag"
	expectedSecond := "Key: 'Config.Files[1].Glob' Error:Field validation for 'Glob' failed on the 'required' tag"
	assert.ErrorContains(t, err, expectedFirst)
	assert.ErrorContains(t, err, expectedSecond)
}

func TestOptionsAreRequired(t *testing.T) {
	b := []byte(`
---	
files:
  - glob: "foo.txt"

`)
	_, err := loadAndValidateConfig(t, b)
	expected := "Key: 'Config.Files[0].Options' Error:Field validation for 'Options' failed on the 'required' tag"
	assert.ErrorContains(t, err, expected)
}

func TestOptionsHitMustBeSuppliedAndPositiveInt(t *testing.T) {
	b := []byte(`
---
files:
  - glob: "foo.txt"
    options:
      hits: -1
`)
	_, err := loadAndValidateConfig(t, b)
	expectedNegative := "Key: 'Config.Files[0].Options.Hits' Error:Field validation for 'Hits' failed on the 'gt' tag"
	assert.ErrorContains(t, err, expectedNegative)
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
	_, err := loadAndValidateConfig(t, b)
	expected := "Key: 'Config.Files[0].Options.Period' Error:Field validation for 'Period' failed on the 'is-valid-time-duration' tag"
	assert.ErrorContains(t, err, expected)
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
	_, err := loadAndValidateConfig(t, b)
	expected := "Key: 'Config.Files[0].Options.Patterns' Error:Field validation for 'Patterns' failed on the 'required' tag"
	assert.ErrorContains(t, err, expected)
}

func TestTriggerMustBeInChoices(t *testing.T) {
	b := []byte(`
---
files:
  - glob: "ok.log"
    options:
      hits: 1
      period: 10s
      patterns:
        - ".*ok"
      trigger: notinlist
`)
	_, err := loadAndValidateConfig(t, b)
	expected := "Key: 'Config.Files[0].Options.Trigger' Error:Field validation for 'Trigger' failed on the"
	assert.ErrorContains(t, err, expected)
}

func TestValidTriggerIsOk(t *testing.T) {
	b := []byte(`
---
files:
  - glob: "ok.log"
    options:
      hits: 1
      period: 10s
      patterns:
        - ".*ok"
      trigger: cloudwatch
`)
	_, err := loadAndValidateConfig(t, b)
	assert.Nil(t, err)
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
