test: pack
	gometalinter.v2 ./...
	go test ./...

pack:
	packr

.PHONY: packr test
