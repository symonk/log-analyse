package config

import (
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	validDuration = "is-valid-time-duration"
)

// isValidTimeDurationFunc allows custom validation on a string type
// that should conform to a time.Duration instance when parsed.
func isValidTimeDurationFunc(fl validator.FieldLevel) bool {
	_, err := time.ParseDuration(fl.Field().String())
	return err == nil
}
