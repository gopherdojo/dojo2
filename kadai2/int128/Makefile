.PHONY: all test doc

all: kadai2

kadai2: *.go */*.go
	go build -o kadai2

test: *.go */*.go
	go test -v ./...

doc: *.go */*.go
	godoc -ex cmd/github.com/gopherdojo/dojo2/kadai2/int128/images
	godoc -ex cmd/github.com/gopherdojo/dojo2/kadai2/int128/options
