package links

import (
	"net/url"
	"regexp"
	"strings"
)

type RegexpLink struct {
	BaseLink
}

func (l RegexpLink) PathString() string {
	return l.Path[1 : len(l.Path)-1]
}

// GetRegexp returns a compiled regexp or an error
func (l RegexpLink) GetRegexp() (*regexp.Regexp, error) {
	return regexp.Compile(l.PathString())
}

func (l RegexpLink) LiteralPrefix() string {
	regexp, err := l.GetRegexp()
	if err != nil {
		return l.Path
	}
	prefix, _ := regexp.LiteralPrefix()
	return prefix
}

func (l RegexpLink) Len() int {
	return len(l.Path) - 2
}

func (l RegexpLink) Match(path string) (bool, error) {
	linkRegExp, err := l.GetRegexp()
	if err != nil {
		return false, err
	}
	return linkRegExp.MatchString(path), nil
}

// Transform returns the target for the given path
func (l RegexpLink) Transform(path string, originalURL url.URL) (*url.URL, error) {
	linkRegExp, err := l.GetRegexp()
	if err != nil {
		return &url.URL{}, err
	}
	return targetToURL(linkRegExp.ReplaceAllString(path, l.Target), originalURL)
}

func (l RegexpLink) MatchPrefix(prefix string) bool {
	// Ideally we would do a regexp intersection here, but we don't have a library for that
	return strings.HasPrefix(l.LiteralPrefix(), prefix) || strings.HasPrefix(prefix, l.LiteralPrefix())
}
