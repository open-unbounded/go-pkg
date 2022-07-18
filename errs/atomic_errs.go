package errs

import (
	"strings"
	"sync"
)

// AtomicErrs an error that can hold multiple errors.
type AtomicErrs struct {
	rw   sync.RWMutex
	errs []error
}

// Add adds errs to be, nil errors are ignored.
func (a *AtomicErrs) Add(errs ...error) {
	a.rw.Lock()
	for _, err := range errs {
		if err == nil {
			continue
		}

		a.errs = append(a.errs, err)
	}
	a.rw.Unlock()
}

// Err returns an error that represents all errors.
func (a *AtomicErrs) Error() string {
	a.rw.RLock()
	defer a.rw.RUnlock()

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
func (a *AtomicErrs) NotNil() bool {
	a.rw.RLock()
	defer a.rw.RUnlock()

	return len(a.errs) > 0
}
