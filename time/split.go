package time

import (
	"math"
	"time"
)

// Split Returns a delimited slice, where the elements of the slice are arrays,
// the first bit of the array is the start time, and the second bit is the end time.
func Split(start, end time.Time, interval time.Duration) [][2]time.Time {
	duration := end.Sub(start)
	if duration <= 0 || interval == 0 {
		return nil
	}

	times := make([][2]time.Time, 0, int(math.Ceil(float64(duration)/float64(interval))))
	p1 := start
	p2 := start.Add(interval)
	// p2 < end => p1 < end
	for ; p2.Before(end); p1, p2 = p2, p2.Add(interval) {
		times = append(times, [2]time.Time{p1, p2})
	}

	times = append(times, [2]time.Time{p1, end})

	return times
}
