.PHONY: all build clean

APP_NAME := server

all: clean build

build: build-linux build-windows build-macos

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)_linux_amd64

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME)_windows_amd64.exe

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)_darwin_amd64

clean:
	rm -rf bin
	mkdir bin
