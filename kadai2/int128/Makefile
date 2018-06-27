.PHONY: all test doc showReaderWriter

all: kadai2

kadai2: *.go */*.go
	go build -o kadai2

test: *.go */*.go
	go test -v -cover ./...

doc: *.go */*.go
	godoc -ex cmd/github.com/gopherdojo/dojo2/kadai2/int128/images
	godoc -ex cmd/github.com/gopherdojo/dojo2/kadai2/int128/options

showReaderWriterImplements:
	cd `go env GOROOT`/src && \
		find . -name '*.go' -and -not -name '*_test.go' -and -not -path '*/internal/*' -and -not -path '*/vendor/*' | \
		xargs egrep -R 'func \(\w+ \*?[A-Z]\w+\) (Read|Write)\(\w+ \[\]byte\)' | \
		column -t -s:

showReaderWriterRefs:
	cd `go env GOROOT`/src && \
		find . -name '*.go' -and -not -name '*_test.go' -and -not -path '*/internal/*' -and -not -path '*/vendor/*' | \
		xargs egrep -R 'func .* io.(Reader|Writer)' | \
		column -t -s:
