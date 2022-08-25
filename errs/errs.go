package errs

import (
	"errors"
	"strings"
)

// Errs an error that can hold multiple errors.
type Errs struct {
	errs []error
}

// Add adds errs to be, nil errors are ignored.
func (a *Errs) Add(errs ...error) {
	for _, err := range errs {
		if err == nil {
			continue
		}

		a.errs = append(a.errs, err)
	}
}

// Err returns an error that represents all errors.
func (a *Errs) Error() string {
	if len(a.errs) == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString(a.errs[0].Error())
	for _, err := range a.errs[1:] {
		builder.WriteRune('\n')
		builder.WriteString(err.Error())
	}

	return builder.String()
}

// NotNil checks if any error inside.
func (a *Errs) NotNil() bool {
	return len(a.errs) > 0
}

// Errors returns a copy of errs.
func (a *Errs) Errors() []error {
	errs := make([]error, len(a.errs))
	copy(errs, a.errs)
	return errs
}

// Is reports whether any errs in err's chain matches target.
func (a *Errs) Is(target error) bool {
	for _, e := range a.errs {
		if errors.Is(e, target) {
			return true
		}
	}

	return false
}

// As finds the first errs in err's chain that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
func (a *Errs) As(target interface{}) bool {
	for _, err := range a.errs {
		if errors.As(err, target) {
			return true
		}
	}

	return false
}
