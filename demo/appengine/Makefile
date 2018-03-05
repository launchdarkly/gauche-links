EXTENSION_NAME = GaucheLinks Demo
PREFIX = go

GCLOUD_BASE=/usr/local/Caskroom/google-cloud-sdk/latest/google-cloud-sdk
APP_SERVER=$(GCLOUD_BASE)/bin/dev_appserver.py

HOST = http://localhost:8080/

ICON_PATH = ../../server/static/images/favicon.png

EXTENSIONS_PATH = build/extensions

# TODO: find a way that works with multiple GOPATH entries
GAUCHE_LINKS = $(GOPATH)/bin/gauche-links

run:
	mkdir -p $(EXTENSIONS_PATH)
	go install github.com/launchdarkly/gauche-links
	$(GAUCHE_LINKS) -extensions -prefix $(PREFIX) -extensions-path $(EXTENSIONS_PATH) -icon $(ICON_PATH) -version 0.0.0.0
	$(APP_SERVER) app.yaml

.PHONY: run
