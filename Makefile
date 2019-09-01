PACKR = $(GOPATH)/bin/packr

init: $(PACKR)

$(PACKR):
	go install github.com/gobuffalo/packr/packr

test: test-extensions
	go test ./...

test-extensions:
	dir="$$(mktemp -d)"; \
	trap 'rm -rf "$${dir}"' EXIT; \
	go run . -extensions -prefix test -extensions-path "$${dir}"; \

packr:
	$(PACKR)

.PHONY: packr test test-extensions