test: pack
	gometalinter.v2 ./...
	go test ./...

pack:
	packr build

.PHONY: packr test
