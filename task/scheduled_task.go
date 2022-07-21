package task

import (
	"fmt"
	"time"

	syncx "github.com/open-unbounded/go-pkg/sync"
	"go.uber.org/atomic"
)

type ScheduledTask struct {
	do          func()
	duration    *atomic.Duration
	startDoOnce func()
	stopDoOnce  func()
	stopChan    chan struct{}
}

func NewScheduledTask(do func(), duration time.Duration) *ScheduledTask {
	s := &ScheduledTask{do: do, duration: atomic.NewDuration(duration), stopChan: make(chan struct{})}
	s.startDoOnce = syncx.DoOnce(s.doStart)
	s.stopDoOnce = syncx.DoOnce(s.doStop)

	return s
}

func (s *ScheduledTask) Start() {
	s.startDoOnce()
}

func (s *ScheduledTask) doStart() {
	go func() {
		s.do()
		ticker := time.NewTicker(s.duration.Load())
		for {
			<-ticker.C
			s.do()

			select {
			case <-s.stopChan:
				fmt.Println("退出")
				return
			default:
			}
			ticker.Reset(s.duration.Load())

		}
	}()
}

func (s *ScheduledTask) Stop() {
	s.stopDoOnce()
}

func (s ScheduledTask) SetDuration(duration time.Duration) {
	s.duration.Store(duration)
}

func (s *ScheduledTask) doStop() {
	close(s.stopChan)
}
