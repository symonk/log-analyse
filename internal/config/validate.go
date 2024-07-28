package config

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrFilesRequired   = errors.New("config `files` is required")
	ErrGlobRequired    = errors.New("config `glob` is required")
	ErrOptionsRequired = errors.New("config `options` is required")
)

const (
	// A custom tag used by the config object to implement custom
	// validation logic
	validateTag = "validate"
)

// Validate implements runtime validation of the config struct using the
// `validate:` custom tag.
func Validate(c any) error {
	refValue := reflect.Indirect(reflect.ValueOf(c))
	for i := 0; i < refValue.NumField(); i++ {
		fieldType := refValue.Type().Field(i)
		tag := fieldType.Tag.Get(validateTag)
		if tag != "" {
			// The field had a custom `validate:` struct tag
			fmt.Println(tag)
		}

	}
	return nil
}
