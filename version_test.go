package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, version, "v0.1.0")
	assert.Equal(t, name, "log-analyse")
}
