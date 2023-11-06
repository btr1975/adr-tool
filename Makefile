# GO-Lang Helpers
# Version: 2023.2.14.001
# Author: Ben Trachtenberg
#
# Description:
#     This is a Makefile to help in building and testing GO projects
#
# Notes:
#     To see all OS/Architecture's you can build use "go tool dist list"
#
#     For compiling for ARM based architecture's you may require the GOARM variable
#     see docs for more info
#
BINARY_NAME=adr-tool
BINARY_DIRECTORY=temp

.PHONY: all test coverage coverage-html format tidy build-windows build-linux build-darwin

all: format tidy coverage-html

test:
	go test ./...

coverage:
	go test ./... -cover

coverage-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

format:
	go fmt ./...

tidy:
	go mod tidy

build-all: build-windows build-linux build-darwin

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_DIRECTORY}/${BINARY_NAME}-windows-amd64.exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_DIRECTORY}/${BINARY_NAME}-linux-amd64

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY_DIRECTORY}/${BINARY_NAME}-darwin-amd64
