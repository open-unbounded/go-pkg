package worker

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/atomic"
)

func TestWorker_Run(t *testing.T) {
	worker := New(100)
	waitGroup := sync.WaitGroup{}
	n := atomic.NewInt32(0)
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		assert.True(t, n.Add(1) <= 101)

		worker.Run(func() {
			defer func() {
				waitGroup.Done()
				n.Sub(1)
			}()
			time.Sleep(time.Millisecond)
		})
	}

	waitGroup.Wait()

}
