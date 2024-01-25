run: 
	go run cmd/main.go

build:
	go build -o bin/item-service cmd/main.go

unit-test:
	go clean -testcache
	go test ./... --cover