package re

import (
	"errors"
	"regexp"
)

// CompileSlice takes a slice of string of what could be
// regular expression patterns, attempts to compile them and return
// the compiled pattern for use along with a slice of errors for any
// that failed.
func CompileSlice(patterns []string) ([]*regexp.Regexp, error) {
	errs := make([]error, 0, len(patterns))
	regexps := make([]*regexp.Regexp, 0, len(patterns))
	for _, pattern := range patterns {
		c, err := regexp.Compile(pattern)
		if err != nil {
			errs = append(errs, err)
		} else {
			regexps = append(regexps, c)
		}
	}
	return regexps, errors.Join(errs...)
}
