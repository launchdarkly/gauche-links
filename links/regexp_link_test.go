package links

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexpLinkMatch(t *testing.T) {
	specs := []struct {
		expectedMatch bool
		path          string
		link          string
	}{
		{true, "same", "same"},
		{false, "same", "different"},
		{true, "same/", "same/"},
		{true, "same/", "same"},
		{false, "same", "same/"},
		{false, "same", "same/"},
		{false, "same/", "same/xyz"},
	}

	for _, s := range specs {
		t.Run(fmt.Sprintf("%s match %s", s.path, s.link), func(t *testing.T) {
			link := NewLink(BaseLink{Path: fmt.Sprintf("/%s/", s.link)})
			actualMatch, err := link.Match(s.path)
			if assert.NoError(t, err) {
				assert.Equal(t, s.expectedMatch, actualMatch)
			}
		})
	}
}

func TestRegexpLinkMatchPrefix(t *testing.T) {
	specs := []struct {
		expectedMatch bool
		path          string
		link          string
	}{
		{true, "same", "same"},
		{false, "same", "different"},
		{true, "same/", "same/"},
		{true, "same/", "same"},
		{true, "same", "same/"},
		{true, "same/", "same/xyz"},
	}

	for _, s := range specs {
		t.Run(fmt.Sprintf("%s match %s", s.path, s.link), func(t *testing.T) {
			link := NewLink(BaseLink{Path: fmt.Sprintf("/%s/", s.link)})
			actualMatch := link.MatchPrefix(s.path)
			assert.Equal(t, s.expectedMatch, actualMatch)
		})
	}
}

func TestRegexpLinkTransform(t *testing.T) {
	specs := []struct {
		path        string
		linkPath    string
		target      string
		originalURL url.URL
		expectedURL string
	}{
		{"same", "same", "http://new/", url.URL{}, "http://new/"},
		{"same", "same", "/new", url.URL{Scheme: "https", Host: "host"}, "https://host/new"},
		{"same/more", "same", "/new", url.URL{Host: "host"}, "//host/new/more"},
		{"same/more", "same/", "/new/", url.URL{Host: "host"}, "//host/new/more"},
		{"same/more", `same/(\w+)`, "/new/$1", url.URL{Host: "host"}, "//host/new/more"},
	}

	for _, s := range specs {
		t.Run(fmt.Sprintf("%s (%s) -> %s", s.path, s.originalURL.String(), s.target), func(t *testing.T) {
			link := NewLink(BaseLink{Path: fmt.Sprintf("/%s/", s.linkPath), Target: s.target})
			actualURL, err := link.Transform(s.path, s.originalURL)
			if assert.NoError(t, err) {
				assert.Equal(t, s.expectedURL, actualURL.String())
			}
		})
	}
}
