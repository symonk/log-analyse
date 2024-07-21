package config

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCanUnmarshalConfigSuccessfully(t *testing.T) {
	config := loadConfigFile(t, "basic.yaml")
	assert.Len(t, config.Files, 2)
}

func TestReturnGlobs(t *testing.T) {
	fConfigs := []FileConfig{{Glob: "foo"}, {Glob: "bar"}, {Glob: "baz"}}
	c := Config{Files: fConfigs}
	assert.Equal(t, c.Globs(), []string{"foo", "bar", "baz"})
}

// loadConfig loads a yaml config file on disk from the /assets/configs
// directory and unmarshals it through viper to return a Config instance
// for testing.
func loadConfigFile(t *testing.T, name string) *Config {
	var c *Config
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	p := path.Join(basepath, "../..", "assets", "configs", name)
	viper.SetConfigFile(p)
	if err := viper.Unmarshal(&c); err != nil {
		t.Fatal("could not load config")
	}
	return c
}
