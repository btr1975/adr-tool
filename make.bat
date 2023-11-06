@ECHO off
REM GO-Lang Helpers
REM Author: Ben Trachtenberg
REM Version: 2023.2.14.001
REM
REM
REM Description:
REM     This is a Makefile to help in building and testing GO projects
REM
REM Notes:
REM     To see all OS/Architecture's you can build use "go tool dist list"
REM
REM     For compiling for ARM based architecture's you may require the GOARM variable
REM     see docs for more info
REM
set BINARY_NAME=adr-tool
set BINARY_DIRECTORY=temp
set WINDOWS_BINARY_DIRECTORY=%BINARY_DIRECTORY%\windows-amd64
set LINUX_BINARY_DIRECTORY=%BINARY_DIRECTORY%\linux-amd64
set DARWIN_BINARY_DIRECTORY=%BINARY_DIRECTORY%\darwin-amd64

IF "%1" == "all" (
    go fmt ./...
    go mod tidy
    go test ./... -v -coverprofile=coverage.out
    go tool cover -html=coverage.out -o coverage.html
    GOTO END
)

IF "%1" == "test" (
    go test ./... %2
    GOTO END
)

IF "%1" == "coverage" (
    go test ./... %2 -cover
    GOTO END
)

IF "%1" == "coverage-html" (
    go test ./... %2 -coverprofile=coverage.out
    go tool cover -html=coverage.out -o coverage.html
    GOTO END
)

IF "%1" == "format" (
    go fmt ./...
    GOTO END
)

IF "%1" == "tidy" (
    go mod tidy
    GOTO END
)

IF "%1" == "build-all" (
    set GOOS=windows
    set GOARCH=amd64
    go build -o %WINDOWS_BINARY_DIRECTORY%\%BINARY_NAME%.exe
    set GOOS=linux
    set GOARCH=amd64
    go build -o %LINUX_BINARY_DIRECTORY%\%BINARY_NAME%
    set GOOS=darwin
    set GOARCH=amd64
    go build -o %DARWIN_BINARY_DIRECTORY%\%BINARY_NAME%
    set GOOS=windows
    set GOARCH=amd64
    GOTO END
)

IF "%1" == "build-windows" (
    set GOOS=windows
    set GOARCH=amd64
    go build -o %WINDOWS_BINARY_DIRECTORY%\%BINARY_NAME%.exe
    GOTO END
)

IF "%1" == "build-linux" (
    set GOOS=linux
    set GOARCH=amd64
    go build -o %LINUX_BINARY_DIRECTORY%\%BINARY_NAME%
    set GOOS=windows
    set GOARCH=amd64
    GOTO END
)

IF "%1" == "build-darwin" (
    set GOOS=darwin
    set GOARCH=amd64
    go build -o %DARWIN_BINARY_DIRECTORY%\%BINARY_NAME%
    set GOOS=windows
    set GOARCH=amd64
    GOTO END
)

:END
