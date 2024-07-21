package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanUnmarshalConfigSuccessfully(t *testing.T) {
	cfg := `
	---
	files:
	# A glob based folder lookup
	- glob: "~/logs/*.txt"
		threshold:
		times: 5
		period: 30s
		patterns:
			- ".*FATAL.*"
			- ".*payment failed.*"
		notify: "email"
	# An explicit log file
	- glob: "~/logs/foo.log"
		threshold:
		times: 1
		period: 1m
		patterns:
		- ".*disk space low.*"
		notify: "slack"
	`
	_ = cfg

}

func TestReturnGlobs(t *testing.T) {
	fConfigs := []FileConfig{{Glob: "foo"}, {Glob: "bar"}, {Glob: "baz"}}
	c := Config{Files: fConfigs}
	assert.Equal(t, c.Globs(), []string{"foo", "bar", "baz"})
}
