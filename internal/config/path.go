package config

import (
	"fmt"
	"os"
	"path"
)

// ConfigDefaultFolder returns the default parent
// folder where config yaml files are looked up when not provided
// by the user.
func ConfigDefaultFolder() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to find the user home directory: %w", err)
	}
	root := path.Join(homeDir, ".log-analyse")
	return root, nil
}
