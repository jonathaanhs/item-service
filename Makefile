run: 
	go run cmd/main.go

build:
	go build -o bin/tokenomy-assessment cmd/main.go

unit-test:
	go clean -testcache
	go test ./... --cover