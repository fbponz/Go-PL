.DEFAULT_GOAL := build
.PHONY: fmt vet build

fmt:
	go fmt ./

vet: fmt
	go vet ./

build: vet
	go build

buildAll:
	go fmt ./
	go vet ./
	go build -o bin/hello hello.go
Exec:
	./bin/hello