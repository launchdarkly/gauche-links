// +build !noappengine

package memcache

import (
	"context"

	"google.golang.org/appengine/memcache"
)

func Get(ctx context.Context, key string) (*memcache.Item, error) {
	return memcache.Get(ctx, key)
}

func Set(ctx context.Context, item *memcache.Item) error {
	return memcache.Set(ctx, item)
}

func Add(ctx context.Context, item *memcache.Item) error {
	return memcache.Add(ctx, item)
}

func Delete(ctx context.Context, key string) error {
	return memcache.Delete(ctx, key)
}

func IsMiss(err error) bool {
	return err == memcache.ErrCacheMiss
}

func NewItem(ctx context.Context) *memcache.Item {
	return new(memcache.Item)
}
