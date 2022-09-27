package strs

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func TestUnderscore(t *testing.T) {
	assert.Equal(t, "", Underscore(""))
	assert.Equal(t, "foo", Underscore("Foo"))
	assert.Equal(t, "foo1", Underscore("Foo1"))
	assert.Equal(t, "foo1", Underscore("foo1"))
	assert.Equal(t, "foo_bar", Underscore("fooBar"))
	assert.Equal(t, "foo_bar", Underscore("fooBar"))
	assert.Equal(t, "foo_bar1", Underscore("fooBar1"))
	assert.Equal(t, "foo1_bar1", Underscore("foo1Bar1"))
	assert.Equal(t, "foo1bar1", Underscore("foo1bar1"))
	assert.Equal(t, "foo1_bar1_bar2", Underscore("foo1Bar1Bar2"))
	assert.Equal(t, "_foo1_bar1_bar2", Underscore("_foo1Bar1Bar2"))
	assert.Equal(t, "_foo1_bar1_bar2", Underscore("_Foo1Bar1Bar2"))
	assert.Equal(t, "___foo1_bar1_bar2", Underscore("___Foo1Bar1Bar2"))
	assert.Equal(t, "__1_foo1_bar1_bar2", Underscore("__1Foo1Bar1Bar2"))
	assert.Equal(t, "__s1_foo1_bar1_bar2", Underscore("__s1Foo1Bar1Bar2"))
}

func TestTruncate(t *testing.T) {
	s := "空山不见人，但闻人语响。返影入深林，复照青苔上。ABCabc,日a本b語ç日ð本Ê語þ日¥本¼語i日©☺☻☹"
	for i := 0; i < 1000; i++ {
		assert.Equal(t, pureTruncateWithTune(s, i), Truncate(s, i))
	}
}

func FuzzTruncate(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string, i int) {
		if utf8.ValidString(s) {
			assert.Equal(t, pureTruncateWithTune(s, i), Truncate(s, i))
		}
	})
}

func truncateWithDecodeRuneInString(s string, length int) string {
	if length == 0 {
		return ""
	}

	if length < 0 {
		return s
	}

	if len(s) <= length {
		return s
	}

	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}

	return s[:n]
}

func truncateWithTune(s string, length int) string {
	if length == 0 {
		return ""
	}

	if length < 0 {
		return s
	}

	if len(s) <= length {
		return s
	}

	if utf8.RuneCountInString(s) <= length {
		return s
	}

	runes := []rune(s)
	return string(runes[:length])
}

func pureTruncateWithTune(s string, length int) string {
	if length < 0 {
		return s
	}

	runes := []rune(s)
	if length > len(runes) {
		length = len(runes)
	}
	return string(runes[:length])
}

func Benchmark_truncateWithTune(b *testing.B) {
	s := "空山不见人，但闻人语响。返影入深林，复照青苔上。"
	size := utf8.RuneCountInString(s)
	for i := 0; i < b.N; i++ {
		for j := 0; j < size+1; j++ {
			truncateWithTune(s, i)
		}
	}
}

func BenchmarkTruncate(b *testing.B) {
	s := "空山不见人，但闻人语响。返影入深林，复照青苔上。"
	size := utf8.RuneCountInString(s)
	for i := 0; i < b.N; i++ {
		for j := 0; j < size+1; j++ {
			Truncate(s, i)
		}
	}
}

func Benchmark_truncateWithDecodeRuneInString(b *testing.B) {
	s := "空山不见人，但闻人语响。返影入深林，复照青苔上。"
	size := utf8.RuneCountInString(s)
	for i := 0; i < b.N; i++ {
		for j := 0; j < size+1; j++ {
			truncateWithDecodeRuneInString(s, i)
		}
	}
}
