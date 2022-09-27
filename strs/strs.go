package strs

import (
	"strings"
	"unicode"
)

// Underscore returns underlined string.
func Underscore(s string) string {
	if s == "" {
		return ""
	}

	builder := strings.Builder{}
	builder.Grow(len(s))

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 && s[i-1] != '_' {
				builder.WriteRune('_')
			}
			builder.WriteRune(unicode.ToLower(r))
			continue
		}

		builder.WriteRune(r)
	}

	return builder.String()
}

// Truncate returns  truncated string of the specified length and
// s consists entirely of valid UTF-8-encoded runes.
func Truncate(s string, length int) string {
	if length == 0 {
		return ""
	}

	if length < 0 {
		return s
	}

	size := len(s)
	if size <= length {
		return s
	}

	var (
		n, i int
		c    rune
	)
	for i, c = range s {
		if n == length-1 {
			break
		}

		n++
	}

	i += len(string(c))
	if i > size {
		i = size
	}

	return s[:i]
}
