// +build !noappengine

package urlfetch

import (
	"context"
	"net/http"

	"google.golang.org/appengine/urlfetch"
)

func Client(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}
