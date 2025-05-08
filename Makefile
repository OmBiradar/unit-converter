.PHONY: fmt vet lint test build build-linux build-windows build-darwin build-arm64

all: fmt vet lint test build

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	# Assuming you have golangci-lint installed
	golangci-lint run ./...

test:
	go test ./... -v

build: build-linux build-windows build-darwin build-arm64

build-linux:
	GOOS=linux GOARCH=amd64 go build -o unitconverter-linux-amd64 main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o unitconverter-windows-amd64.exe main.go

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o unitconverter-darwin-amd64 main.go

build-arm64:
	GOOS=linux GOARCH=arm64 go build -o unitconverter-linux-arm64 main.go
