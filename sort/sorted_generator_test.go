package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedGenerator_Gen(t *testing.T) {
	type u struct {
		index int
	}

	sg := SortedGenerator[u]{}
	for i := 0; i < 100; i++ {
		assert.EqualValues(
			t,
			u{index: i},
			sg.Gen(func(index int) u {
				return u{index: index}
			}),
		)
	}
}

func TestAtomicSortedGenerator_Gen(t *testing.T) {
	type u struct {
		index int
	}

	sg := AtomicSortedGenerator[u]{}
	for i := 0; i < 100; i++ {
		assert.EqualValues(
			t,
			u{index: i},
			sg.Gen(func(index int64) u {
				return u{index: int(index)}
			}),
		)
	}
}
