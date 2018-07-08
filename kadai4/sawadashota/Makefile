dev:
	go run main.go
test:
	go test -v ./...
coverage:
	go test -coverprofile=profile ./... && go tool cover -html=profile -o profile.html
