package sync

import "sync"

// DoOne returns a function that will only be executed once.
func DoOne(do func()) func() {
	var once sync.Once
	return func() {
		once.Do(do)
	}
}
