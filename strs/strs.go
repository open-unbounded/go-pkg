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
