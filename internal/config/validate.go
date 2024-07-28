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

func Validate(c *Config) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.RegisterValidation(validDuration, isValidTimeDurationFunc); err != nil {
		return err
	}
	if err := validate.Struct(c); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		return err.(validator.ValidationErrors)

	}
	return nil
}
