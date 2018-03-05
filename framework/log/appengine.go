// +build !noappengine

package log

import (
	"context"

	"google.golang.org/appengine/log"
)

func Debugf(ctx context.Context, format string, args ...interface{}) {
	log.Debugf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	log.Debugf(ctx, format, args...)
}

func Warningf(ctx context.Context, format string, args ...interface{}) {
	log.Warningf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(ctx, format, args...)
}

func Criticalf(ctx context.Context, format string, args ...interface{}) {
	log.Criticalf(ctx, format, args...)
}
