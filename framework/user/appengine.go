// +build !noappengine

package user

import (
	"context"

	"google.golang.org/appengine/user"
)

func Current(ctx context.Context) *user.User {
	return user.Current(ctx)
}
