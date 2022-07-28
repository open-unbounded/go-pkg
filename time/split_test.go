package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {

	start := time.Now()
	end := start.Add(time.Hour - time.Second*40)
	assert.Len(t, Split(start, end, time.Minute), 60)

	start = time.Now()
	end = start.Add(time.Hour + time.Second*40)
	assert.Len(t, Split(start, end, time.Minute), 61)

	start = time.Now()
	end = start.Add(time.Hour)
	assert.Len(t, Split(start, end, time.Minute), 60)

	// ------

	start = time.Now()
	end = start.Add(time.Hour - time.Second*40)
	assert.Len(t, Split(start, end, time.Second), 60*60-40)

	start = time.Now()
	end = start.Add(time.Hour + time.Second*40)
	assert.Len(t, Split(start, end, time.Second), 60*60+40)

	start = time.Now()
	end = start.Add(time.Hour)
	assert.Len(t, Split(start, end, time.Second), 60*60)

	// ---

	start = time.Now()
	end = start
	assert.Len(t, Split(start, end, time.Minute), 0)

	start = time.Now()
	end = start.Add(time.Hour)
	assert.Len(t, Split(start, end, 0), 0)

	// ---

	start = time.Now()
	end = start.Add(time.Hour)
	assert.Len(t, Split(start, end, time.Hour), 1)

}
