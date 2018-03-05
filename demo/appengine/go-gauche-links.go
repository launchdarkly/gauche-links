package appengine

import (
	"github.com/launchdarkly/gauche-links/server"
)

func init() {
	if err := server.Start(); err != nil {
		panic(err)
	}
}
