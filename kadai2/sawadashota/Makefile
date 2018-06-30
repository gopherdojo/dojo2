test:
	go test -v ./...
coverage:
	go test -coverprofile=profile ./... && go tool cover -html=profile -o profile.html
jpeg-to-png:
	go run ./cmd/main.go ./testdata/
png-to-gif:
	go run ./cmd/main.go -f png -t gif ./testdata/
