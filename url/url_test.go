package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	url, err := Url("https://a.com", "a")
	assert.NoError(t, err)
	assert.Equal(t, "https://a.com/a", url)

	url, err = Url("https://a.com", "/a/", "b")
	assert.NoError(t, err)
	assert.Equal(t, "https://a.com/a/b", url)

	url, err = Url("https://a.com", "a", "/b/c")
	assert.NoError(t, err)
	assert.Equal(t, "https://a.com/a/b/c", url)

	url, err = Url("https://a.com/x", "/a/x/", "b/c/")
	assert.NoError(t, err)
	assert.Equal(t, "https://a.com/x/a/x/b/c", url)

	url, err = Url("https://a.com/x/x", "a", "/b/c")
	assert.NoError(t, err)
	assert.Equal(t, "https://a.com/x/x/a/b/c", url)

	url, err = Url("https://a.com/x/x?x=1", "/a", "b/c")
	assert.NoError(t, err)
	assert.Equal(t, "https://a.com/x/x/a/b/c", url)

	url, err = Url("%gh&%ij", "a", "b/c")
	assert.Error(t, err)
	assert.Equal(t, "", url)
}
