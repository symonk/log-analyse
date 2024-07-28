package config

import "errors"

var (
	ErrFilesRequired   = errors.New("config `files` is required")
	ErrGlobRequired    = errors.New("config `glob` is required")
	ErrOptionsRequired = errors.New("config `options` is required")
)
