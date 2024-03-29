package strs

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoiner_WriteString(t *testing.T) {
	join := NewJoiner(WithJoiner("(", ",", ")"))
	_, _ = join.WriteString("1")
	_, _ = join.WriteString("2")
	_, _ = join.WriteString("3")
	assert.Equal(t, "(1,2,3)", join.String())

	join = NewJoiner(WithJoiner("(", ",", ""))
	_, _ = join.WriteString("1")
	_, _ = join.WriteString("2")
	_, _ = join.WriteString("3")
	assert.Equal(t, "(1,2,3", join.String())

	join = NewJoiner(WithJoinerStep("-"))
	_, _ = join.WriteString("1")
	_, _ = join.WriteString("2")
	_, _ = join.WriteString("3")
	assert.Equal(t, "1-2-3", join.String())

	join = NewJoiner(WithJoinerStep("-"), WithJoinerPrefix("=>"))
	_, _ = join.WriteString("1")
	_, _ = join.WriteString("2")
	_, _ = join.WriteString("3")
	assert.Equal(t, "=>1-2-3", join.String())

	join = NewJoiner(WithJoinerStep("-"), WithJoinerPrefix("=>"), WithJoinerSuffix("<===="))
	_, _ = join.WriteString("1")
	_, _ = join.WriteString("2")
	_, _ = join.WriteString("3")
	assert.Equal(t, "=>1-2-3<====", join.String())
}

func TestJoiner_Write(t *testing.T) {
	join := NewJoiner(WithJoiner("(", ",", ")"))
	_ = join.WriteByte('a')
	_ = join.WriteByte('b')
	_ = join.WriteByte('c')
	assert.Equal(t, "(a,b,c)", join.String())

	join = NewJoiner(WithJoiner("(", ",", ""))
	_ = join.WriteByte('a')
	_ = join.WriteByte('b')
	_ = join.WriteByte('c')
	assert.Equal(t, "(a,b,c", join.String())

	join = NewJoiner(WithJoinerStep("000"))
	_ = join.WriteByte('a')
	_ = join.WriteByte('b')
	_ = join.WriteByte('c')
	assert.Equal(t, "a000b000c", join.String())

	join = NewJoiner(WithJoinerStep("1"), WithJoinerPrefix("--"))
	_ = join.WriteByte('a')
	_ = join.WriteByte('b')
	_ = join.WriteByte('c')
	assert.Equal(t, "--a1b1c", join.String())

	join = NewJoiner(WithJoinerStep("1"), WithJoinerPrefix("--"), WithJoinerSuffix("<---"))
	_ = join.WriteByte('a')
	_ = join.WriteByte('b')
	_ = join.WriteByte('c')
	assert.Equal(t, "--a1b1c<---", join.String())
}

func TestJoiner_WriteRune(t *testing.T) {
	join := NewJoiner(WithJoiner("(", ",", ")"))
	_, _ = join.WriteRune('a')
	_, _ = join.WriteRune('b')
	_, _ = join.WriteRune('c')
	assert.Equal(t, "(a,b,c)", join.String())

	join = NewJoiner(WithJoiner("(", ",", ""))
	_, _ = join.WriteRune('a')
	_, _ = join.WriteRune('b')
	_, _ = join.WriteRune('c')
	assert.Equal(t, "(a,b,c", join.String())

	join = NewJoiner(WithJoinerStep("000"))
	_, _ = join.WriteRune('a')
	_, _ = join.WriteRune('b')
	_, _ = join.WriteRune('c')
	assert.Equal(t, "a000b000c", join.String())

	join = NewJoiner(WithJoinerStep("1"), WithJoinerPrefix("--"))
	_ = join.WriteByte('a')
	_ = join.WriteByte('b')
	_ = join.WriteByte('c')
	assert.Equal(t, "--a1b1c", join.String())

	join = NewJoiner(WithJoinerStep("1"), WithJoinerPrefix("--"), WithJoinerSuffix("<---"))
	_ = join.WriteByte('a')
	_ = join.WriteByte('b')
	_ = join.WriteByte('c')
	assert.Equal(t, "--a1b1c<---", join.String())
}

func TestJoiner_Len(t *testing.T) {
	join := NewJoiner(WithJoiner("(", ",", ")"))
	assert.Equal(t, 2, join.Len())
	_, _ = join.WriteRune('a')
	_, _ = join.WriteRune('b')
	_, _ = join.WriteRune('c')
	assert.Equal(t, 7, join.Len())
	assert.Equal(t, len(join.String()), join.Len())
}

func TestJoiner_Grow(t *testing.T) {
	for _, growLen := range []int{0, 100, 1000, 10000, 100000} {
		p := bytes.Repeat([]byte{'a'}, growLen)
		b := NewJoiner()
		allocs := testing.AllocsPerRun(100, func() {
			b.Reset()
			b.Grow(growLen) // should be only alloc, when growLen > 0
			if b.Cap() < growLen {
				t.Fatalf("growLen=%d: Cap() is lower than growLen", growLen)
			}
			_, _ = b.Write(p)
			if b.String() != string(p) {
				fmt.Println(b.String(), "  ", string(p))
				fmt.Println(len(b.String()), "  ", len(string(p)))
				t.Fatalf("growLen=%d: bad data written after Grow", growLen)
			}
		})
		wantAllocs := 1
		if growLen == 0 {
			wantAllocs = 0
		}
		if g, w := int(allocs), wantAllocs; g != w {
			t.Errorf("growLen=%d: got %d allocs during Write; want %v", growLen, g, w)
		}
	}
}

func BenchmarkNewJoin(b *testing.B) {
	join := NewJoiner(WithJoiner("(", ",", ")"))
	for i := 0; i < b.N; i++ {
		_, _ = join.WriteString("1")
		_, _ = join.WriteString("2")
		_, _ = join.WriteString("3")
	}
}

func TestJoiner_Cap(t *testing.T) {
	join := NewJoiner()
	assert.Equal(t, 0, join.Cap())
	_, _ = join.WriteString("1")
	assert.Equal(t, 8, join.Cap())
}

func TestJoiner_Reset(t *testing.T) {
	join := NewJoiner()
	assert.Equal(t, 0, join.Cap())
	_, _ = join.WriteString("111")
	assert.Equal(t, 8, join.Cap())
	join.Reset()
	assert.Equal(t, 0, join.Cap())
}
