package git

import (
	"errors"
	"net/url"
	"path"
	"regexp"
	"strings"
)

// URL is an interface that represents a git url
type URL interface {
	ToGoPath() string
}

// ErrInvalidURL is returned from Parse when the url is not a valid git url
var ErrInvalidURL = errors.New("Invalid URL")

var gitURLParser = regexp.MustCompile("^git@([\\w\\.]+):(.+?)(?:\\.git)?$")
var gitPathParser = regexp.MustCompile("^(.+?)(?:\\.git)?$")

// Parse gets a url string and returns a URL object, or an error
func Parse(uri string) (URL, error) {
	if uri == "" {
		return nil, ErrInvalidURL
	}
	if strings.HasPrefix(uri, "git@") {
		if !gitURLParser.MatchString(uri) {
			return nil, ErrInvalidURL
		}

		matches := gitURLParser.FindStringSubmatch(uri)
		if len(matches) != 3 {
			return nil, ErrInvalidURL
		}

		return gitURL{
			domain: matches[1],
			path:   matches[2],
		}, nil
	}
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		u, err := url.Parse(uri)
		if err != nil {
			return gitURL{}, nil
		}
		matches := gitPathParser.FindStringSubmatch(u.Path)
		if len(matches) != 2 {
			return nil, ErrInvalidURL
		}

		return gitURL{
			domain: u.Hostname(),
			path:   matches[1],
		}, nil
	}
	return gitURL{}, nil
}

type gitURL struct {
	domain string
	path   string
}

func (g gitURL) ToGoPath() string {
	return path.Join(g.domain, g.path)
}
