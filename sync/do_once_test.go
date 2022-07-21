package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoOne(t *testing.T) {
	i := 1
	doOnce := DoOnce(func() {
		i++
	})
	assert.Equal(t, 1, i)
	doOnce()
	assert.Equal(t, 2, i)
	doOnce()
	assert.Equal(t, 2, i)
}
