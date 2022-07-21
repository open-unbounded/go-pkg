package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoOne(t *testing.T) {
	i := 1
	doOne := DoOne(func() {
		i++
	})
	assert.Equal(t, 1, i)
	doOne()
	assert.Equal(t, 2, i)
	doOne()
	assert.Equal(t, 2, i)
}
