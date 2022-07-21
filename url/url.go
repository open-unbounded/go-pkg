package url

import (
	"net/url"
	"path"
)

func Url(addr string, paths ...string) (string, error) {
	nurl, err := url.Parse(addr)
	if err != nil {
		return "", err
	}

	return nurl.Scheme + "://" + path.Join(append([]string{nurl.Host, nurl.Path}, paths...)...), nil
}
