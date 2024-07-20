package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultBaseDirSuccess(t *testing.T) {
	home, _ := os.UserHomeDir()
	expected := fmt.Sprintf("%s/%s", home, ".log-analyse")
	baseDir, err := ConfigDefaultFolder()
	assert.Nil(t, err)
	assert.Equal(t, baseDir, expected)
}
