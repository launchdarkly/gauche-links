package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/launchdarkly/gauche-links/extension"
)

func main() {
	var version string
	var prefix string
	var host string
	var devHost string
	var extensionsPath string
	var icon string
	var name string
	var platform string

	flag.StringVar(&extensionsPath, "extensions-path", ".", "Path to place the extensions")
	flag.StringVar(&icon, "icon", "", "Extension icon")
	flag.StringVar(&version, "version", "0.0.0.0", "Extension version")
	flag.StringVar(&host, "host", "http://localhost:8080", "Production host url")
	flag.StringVar(&devHost, "dev-host", "http://localhost:8080", "Development host url")
	flag.StringVar(&prefix, "prefix", "go", "Link prefix")
	flag.StringVar(&name, "name", "GaucheLinks", "Extension name")
	flag.StringVar(&platform, "extension", "", `Extension platform when generating extensions ("chrome" or "firefox")`)

	flag.Parse()

	if platform != "" {
		config := extension.Config{
			Prefix:   prefix,
			Version:  version,
			Host:     host,
			DevHost:  devHost,
			Icon:     icon,
			Name:     name,
			Platform: platform,
		}
		if err := extension.Build(config, extensionsPath); err != nil {
			fmt.Fprintf(os.Stderr, "unable to build extensions: %s\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Fprintf(os.Stderr, "server is not implemented")
	os.Exit(1)
}
