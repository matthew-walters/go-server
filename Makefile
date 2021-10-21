fmt:
	@go fmt ./...

vet:
	@go vet ./...

imports:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -v -l -w .

build:
	go build -o bin/main main.go

run:
	go run main.go