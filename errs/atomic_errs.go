package errs

import (
	"strings"
	"sync"
)

type AtomicErrs struct {
	rw   sync.RWMutex
	errs []error
}

func (a *AtomicErrs) Add(err error) {
	a.rw.Lock()
	a.errs = append(a.errs, err)
	a.rw.Unlock()
}

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
