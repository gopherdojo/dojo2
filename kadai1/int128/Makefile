.PHONY: all test doc

all: kadai1

kadai1: *.go */*.go
	go build -o kadai1

test: *.go */*.go
	go test -v ./...

doc: *.go */*.go
	godoc -ex cmd/github.com/gopherdojo/dojo2/kadai1/int128/images
	godoc -ex cmd/github.com/gopherdojo/dojo2/kadai1/int128/options
