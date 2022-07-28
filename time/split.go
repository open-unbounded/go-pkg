package time

import (
	"time"
)

func Split(start, end time.Time, interval time.Duration) [][2]time.Time {
	if start.Equal(end) || interval == 0 {
		return nil
	}

	times := make([][2]time.Time, 0, end.Sub(start)/interval)
	p1 := start
	p2 := start.Add(interval)
	// p2 < end => p1 < end
	for ; p2.Before(end); p1, p2 = p2, p2.Add(interval) {
		times = append(times, [2]time.Time{p1, p2})
	}

	times = append(times, [2]time.Time{p1, end})

	return times
}
