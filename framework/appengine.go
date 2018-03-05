// +build !noappengine

package framework

import (
	"context"
	"net/http"

	"google.golang.org/appengine"
)

func NewContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}
