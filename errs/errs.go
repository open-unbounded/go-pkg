package errs

import (
	"strings"
)

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
