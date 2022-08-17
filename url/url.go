package url

import (
	"net/url"
)

// Url returns a URL string with the provided path elements joined to
// the existing path of base and the resulting path cleaned of any ./ or ../ elements.
// Deprecated: Use url.JoinPath instead.
func Url(addr string, paths ...string) (string, error) {
	return url.JoinPath(addr, paths...)
}
