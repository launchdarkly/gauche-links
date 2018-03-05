package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/launchdarkly/gauche-links/extension"
)

func main() {
	var extensions bool
	var version string
	var prefix string
	var host string
	var devHost string
	var extensionsPath string
	var icon string
	var name string

	flag.BoolVar(&extensions, "extensions", false, "Set to true to build extensions")
	flag.StringVar(&extensionsPath, "extensions-path", ".", "Path to place the extensions")
	flag.StringVar(&icon, "icon", "", "Extension icon")
	flag.StringVar(&version, "version", "0.0.0.0", "Extension version")
	flag.StringVar(&host, "host", "http://localhost:8080", "Production host url")
	flag.StringVar(&devHost, "dev-host", "http://localhost:8080", "Development host url")
	flag.StringVar(&prefix, "prefix", "go", "Link prefix")
	flag.StringVar(&prefix, "name", "GaucheLinks", "Tool name")

	flag.Parse()

	if extensions {
		config := extension.ExtensionConfig{
			Prefix:  prefix,
			Version: version,
			Host:    host,
			DevHost: devHost,
			Icon:    icon,
			Name:    name,
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
