.PHONY: all

setup:
	go get -u github.com/kisielk/errcheck

build:
	go build cli.go conversion-cli

test:
	go test ./... -cover

errcheck:
	errcheck ./...

