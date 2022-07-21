package url

import (
	"net/url"
	"path"
)

func Url(add string, paths ...string) (string, error) {
	nurl, err := url.Parse(add)
	if err != nil {
		return "", err
	}

	return nurl.Scheme + "://" + path.Join(append([]string{nurl.Host, nurl.Path}, paths...)...), nil
}
