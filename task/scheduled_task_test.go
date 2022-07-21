package task

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/atomic"
)

func TestScheduledTask(t *testing.T) {
	i := &atomic.Int32{}
	task := NewScheduledTask(func() {
		i.Add(1)
	}, time.Second)
	task.Start()
	defer task.Stop()
	time.Sleep(time.Second + time.Second/2)
	task.Stop()

	time.Sleep(time.Second * 2)
	assert.EqualValues(t, 3, i.Load())
}
