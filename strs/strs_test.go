package strs

import (
	"testing"

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
