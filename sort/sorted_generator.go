package sort

import (
	"sync/atomic"
)

type SortedGenerator[T any] struct {
	index int
}

func (sg *SortedGenerator[T]) Gen(fn func(index int) T) T {
	defer func() { sg.index++ }()
	return fn(sg.index)
}

// -----------------

type AtomicSortedGenerator[T any] struct {
	index atomic.Int64
}

func (sg *AtomicSortedGenerator[T]) Gen(fn func(index int64) T) T {
	defer sg.index.Add(1)
	return fn(sg.index.Load())
}
