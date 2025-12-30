.DEFAULT_GOAL := build
.PHONY: fmt vet build

clean:
	rm -rf ./bin 

fmt: clean
	go fmt ./

vet: fmt
	go vet ./

build: vet
	go build -o bin/test ./*.go


Exec:
	./bin/test