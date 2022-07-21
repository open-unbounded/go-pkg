package sync

import "sync"

// DoOnce returns a function that will only be executed once.
func DoOnce(do func()) func() {
	var once sync.Once
	return func() {
		once.Do(do)
	}
}
