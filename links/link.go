package links

import (
	"net/url"
	"strings"
)

type Link interface {
	Base() BaseLink
	LiteralPrefix() string
	Len() int
	Match(string) (bool, error)
	MatchDescription(string) bool
	MatchPrefix(string) bool
	PathString() string
	Transform(string, url.URL) (*url.URL, error)
}

type LinkImpl struct {
	Link
}

// Link defines each individual link
type BaseLink struct {
	Path        string
	Target      string
	Description string
	Author      string
	EditURL     *url.URL
}

func (l BaseLink) PathString() string {
	return l.Path
}

func (l BaseLink) LiteralPrefix() string {
	return l.Path
}

func (l BaseLink) Base() BaseLink {
	return l
}

func (l BaseLink) Len() int {
	return len(l.Path)
}

// IsRegexp returns whether the link is a regexp match
func (l BaseLink) IsRegexp() bool {
	return strings.HasPrefix(l.Path, "/") && strings.HasSuffix(l.Path, "/")
}

// Match returns whether a link completely starts a string or includes the link as part of it's path
func (l BaseLink) Match(path string) (bool, error) {
	if !strings.HasPrefix(path, l.Path) {
		return false, nil
	}
	return path == l.Path || strings.HasSuffix(l.Path, "/") || path[len(l.Path)] == '/', nil
}

// MatchDescription returns whether a string matches the description
func (l BaseLink) MatchDescription(path string) bool {
	return strings.Contains(strings.ToLower(l.Description), strings.ToLower(path))
}

func targetToURL(target string, originalURL url.URL) (destURL *url.URL, err error) {
	if strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "http://") {
		if destURL, err = url.Parse(target); err != nil {
			return nil, err
		}
	} else if strings.HasPrefix(target, "/") {
		if destURL, err = url.Parse(target); err != nil {
			return nil, err
		}
	} else {
		if destURL, err = url.Parse("https://" + target); err != nil {
			return nil, err
		}
	}

	if !destURL.IsAbs() {
		destURL.Host = originalURL.Host
		destURL.Scheme = originalURL.Scheme
	} else if destURL.Scheme == "" {
		destURL.Scheme = "https"
	}
	if destURL.RawQuery == "" {
		destURL.RawQuery = originalURL.RawQuery
	}
	if destURL.Fragment == "" {
		destURL.Fragment = originalURL.Fragment
	}
	return destURL, nil
}

// Transform returns the target for the given path, appending the rest of the path
func (l BaseLink) Transform(path string, originalURL url.URL) (*url.URL, error) {
	destURL := l.Target
	if len(path) > len(l.Path) {
		destURL += path[len(l.Path):]
	}
	return targetToURL(destURL, originalURL)
}

// MatchPrefix returns a suggestion if given path is a prefix for a link
func (l BaseLink) MatchPrefix(prefix string) bool {
	return strings.HasPrefix(l.Path, prefix)
}

func NewLink(link BaseLink) Link {
	if link.IsRegexp() {
		return RegexpLink{link}
	} else {
		return link
	}
}
