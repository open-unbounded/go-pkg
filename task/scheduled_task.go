package task

import (
	"time"

	syncx "github.com/open-unbounded/go-pkg/sync"
	"go.uber.org/atomic"
)

// ScheduledTask represents a scheduled task.
type ScheduledTask struct {
	do          func()
	duration    *atomic.Duration
	startDoOnce func()
	stopDoOnce  func()
	stopChan    chan struct{}
}

// NewScheduledTask returns a ScheduledTask.
func NewScheduledTask(do func(), duration time.Duration) *ScheduledTask {
	s := &ScheduledTask{do: do, duration: atomic.NewDuration(duration), stopChan: make(chan struct{})}
	s.startDoOnce = syncx.DoOnce(s.doStart)
	s.stopDoOnce = syncx.DoOnce(s.doStop)

	return s
}

// Start starts scheduling tasks.
func (s *ScheduledTask) Start() {
	s.startDoOnce()
}

func (s *ScheduledTask) doStart() {
	s.do()
	ticker := time.NewTicker(s.duration.Load())
	for {
		<-ticker.C
		s.do()

		select {
		case <-s.stopChan:
			return
		default:
		}
		ticker.Reset(s.duration.Load())
	}
}

// Stop gracefully stop a scheduled task.
func (s *ScheduledTask) Stop() {
	s.stopDoOnce()
}

func (s *ScheduledTask) doStop() {
	close(s.stopChan)
}

// SetDuration sets the interval for scheduling tasks.
func (s ScheduledTask) SetDuration(duration time.Duration) {
	s.duration.Store(duration)
}
