package server

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"sort"
	"strings"
	textTemplate "text/template"
	"time"

	"github.com/fabioberger/airtable-go"
	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"

	"github.com/launchdarkly/gauche-links/framework"
	"github.com/launchdarkly/gauche-links/framework/log"
	"github.com/launchdarkly/gauche-links/framework/memcache"
	"github.com/launchdarkly/gauche-links/framework/urlfetch"
	"github.com/launchdarkly/gauche-links/framework/user"
	"github.com/launchdarkly/gauche-links/links"
)

// AirTable defines AirTable config section
type AirTable struct {
	APIKey   string `required:"true"`
	BaseID   string `required:"true"`
	RootURL  string `required:"true"`
	AppPath  string `required:"true"`
	BasePath string `required:"true"`
}

// Config defines global configuration
type Config struct {
	Global
	AirTable
}

type Global struct {
	Prefix string `default:"go"`
	Host   string
}

// LinkCacheKey is the key for the links items in the cache
var LinkCacheKey = "links"

// CacheExpirationSeconds is the umber of seconds to keep data in memcache
var CacheExpirationSeconds = 2

var rootTemplate *template.Template
var errorTemplate *template.Template
var searchSpecTemplate *textTemplate.Template

func init() {
	box := packr.NewBox("./static/templates")
	rootTemplate = template.Must(template.New("index").Parse(box.String("index.html")))
	errorTemplate = template.Must(template.New("error").Parse(box.String("error.html")))
	searchSpecTemplate = textTemplate.Must(textTemplate.New("searchSpec").Parse(box.String("search.xml")))

	if err := start(); err != nil {
		panic(err)
	}
}

var config Config

func start() error {
	if err := envconfig.Process("gauche", &config.Global); err != nil {
		return fmt.Errorf("unable to process environment variables: %s", err)
	}
	if err := envconfig.Process("gauche_airtable", &config.AirTable); err != nil {
		return fmt.Errorf("unable to process environment variables: %s", err)
	}

	r := mux.NewRouter()
	r.Use(requireUser)
	r.HandleFunc("/_search", suggest).Methods("GET")
	r.HandleFunc("/_search.xml", searchSpec).Methods("GET")
	r.HandleFunc("/_logout", logout).Methods("GET")
	r.PathPrefix("/").HandlerFunc(link).Methods("GET")

	http.Handle("/", r)
	return nil
}

// LinkRecord represents a single record from AirTable
type LinkRecord struct {
	AirtableID string `json:"id,omitempty"`
	Fields     struct {
		Path        string
		Target      string
		Description string
		Author      string
	} `json:"fields"`
}

func getLinks(ctx context.Context) ([]links.Link, error) {
	// First, check memcache
	cachedLinks, err := memcache.Get(ctx, LinkCacheKey)
	if err == nil && cachedLinks != nil {
		var baseLinks []links.BaseLink
		unmarshalErr := json.Unmarshal(cachedLinks.Value, &baseLinks)
		if unmarshalErr == nil {
			allLinks := make([]links.Link, len(baseLinks))
			for i, l := range baseLinks {
				allLinks[i] = links.NewLink(l)
			}
			return allLinks, nil
		}
		log.Warningf(ctx, "Unable to decode links from cache: %s", err)
	} else if err != nil && memcache.IsMiss(err) {
		log.Warningf(ctx, "Unable to fetch links from cache: %s", err)
	}

	client, err := airtable.New(config.AirTable.APIKey, config.AirTable.BaseID)
	if err != nil {
		return nil, err
	}
	client.HTTPClient = urlfetch.Client(ctx)

	var records []LinkRecord
	if err = client.ListRecords("Links", &records, airtable.ListParameters{View: "Main"}); err != nil {
		return nil, err
	}

	rootURL, parseErr := url.Parse(config.AirTable.RootURL)
	if parseErr != nil {
		panic(parseErr)
	}

	var allLinks []links.Link
	for _, record := range records {
		if record.Fields.Path != "" {
			editURL := *rootURL
			editURL.Path = config.AirTable.BasePath + "/" + record.AirtableID
			allLinks = append(allLinks, links.NewLink(links.BaseLink{
				Path:        record.Fields.Path,
				Description: record.Fields.Description,
				Target:      record.Fields.Target,
				Author:      record.Fields.Author,
				EditURL:     &editURL,
			}))
		}
	}

	linkData, err := json.Marshal(allLinks)
	if err != nil {
		return nil, err
	}
	cachedLinks = memcache.NewItem(ctx)
	cachedLinks.Key = LinkCacheKey
	cachedLinks.Value = linkData
	cachedLinks.Expiration = time.Duration(CacheExpirationSeconds) * time.Second
	err = memcache.Add(ctx, cachedLinks)
	if err != nil {
		log.Warningf(ctx, "Unable to save links to cache: %s", err)
	}

	return allLinks, nil
}

func link(w http.ResponseWriter, r *http.Request) {
	ctx := framework.NewContext(r)

	var destURL *url.URL
	allLinks, err := getLinks(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to fetch links: %s", err), http.StatusInternalServerError)
	}

	// Remove zero width unicode whitespace (caused by a hack we do in the extension)
	path := strings.Trim(r.URL.Path[1:], " \u200B")

	isEdit := false
	if strings.HasSuffix(path, "/edit") && path != "/edit" {
		isEdit = true
		path = path[0 : len(path)-len("/edit")]
	}

	if path != "" {
		for _, l := range allLinks {
			match, err := l.Match(path)
			if err != nil {
				log.Warningf(ctx, `Unable to compile link "%s"`, l.PathString())
				continue
			}

			if !match {
				continue
			}

			if isEdit {
				base := l.Base()
				http.Redirect(w, r, base.EditURL.String(), http.StatusTemporaryRedirect)
				return
			}

			destURL, err = l.Transform(path, *r.URL)
			if err != nil {
				log.Warningf(ctx, `Bad url "%s" for link "%s": %s`, l.Base().Target, l.Base().Path, err)
				if tmplErr := errorTemplate.Execute(w, map[string]interface{}{"link": l, "err": err}); tmplErr != nil {
					http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			break
		}
	}

	if destURL != nil {
		log.Infof(ctx, `Redirecting to %s`, destURL.String())
		http.Redirect(w, r, destURL.String(), http.StatusTemporaryRedirect)
	}

	var possibleLinks []links.Link
	if path != "" {
		for _, l := range allLinks {
			if l.MatchPrefix(path) || l.MatchDescription(path) {
				possibleLinks = append(possibleLinks, l)
			}
		}
		sort.Sort(links.LinkSliceByPath(links.LinkSliceByPathLength(allLinks)))
	} else {
		possibleLinks = allLinks
		sort.Sort(links.LinkSliceByPath(allLinks))
	}

	u := user.Current(ctx)
	templateData := map[string]interface{}{
		"links":       possibleLinks,
		"prefix":      config.Prefix,
		"query":       path,
		"author":      u.Email,
		"host":        config.Host,
		"originalURL": r.URL,
	}

	if err := rootTemplate.Execute(w, templateData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if path != "" {
		w.WriteHeader(http.StatusNotFound)
	}
}

func suggest(w http.ResponseWriter, r *http.Request) {
	ctx := framework.NewContext(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	allLinks, err := getLinks(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to fetch links: %s", err), http.StatusInternalServerError)
	}

	query := r.URL.Query().Get("q")

	w.Header().Set("Content-Type", "application/json")

	var completions, descriptions, urls []string

	if query != "" {
		for _, l := range allLinks {
			if l.MatchPrefix(query) || l.MatchDescription(query) {
				completions = append(completions, l.Base().Path)
				descriptions = append(descriptions, l.Base().Description)
				urls = append(urls, l.Base().Target)
			}
		}
	}

	data, err := json.Marshal([]interface{}{query, completions, descriptions, urls})
	if err != nil {
		panic(err)
	}
	_, writeErr := w.Write(data)
	if writeErr != nil {
		log.Errorf(ctx, "Unable to write response: %s", writeErr)
	}
}

func searchSpec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")

	templateData := map[string]interface{}{
		"host": config.Host,
	}

	if tmplErr := searchSpecTemplate.Execute(w, templateData); tmplErr != nil {
		http.Error(w, tmplErr.Error(), http.StatusInternalServerError)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	ctx := framework.NewContext(r)
	logoutURL, err := user.LogoutURL(ctx, "/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, logoutURL, http.StatusTemporaryRedirect)
}

func requireUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := framework.NewContext(r)
		if user.Current(ctx) == nil {
			http.Error(w, "expected user to be logged in", http.StatusInternalServerError)
			return
		}
		h.ServeHTTP(w, r)
	})
}
