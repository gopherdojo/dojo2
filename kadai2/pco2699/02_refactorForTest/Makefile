.PHONY: all test cover

all: kadai2

kadai2: *.go */*.go
		go build -o kadai2

test: *.go */*.go
		go test -v ./lib

cover: *.go */*.go
		go test -coverprofile=cover.out ./lib
		go tool cover -html=cover.out