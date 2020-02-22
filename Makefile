all: clean build test

clean:
	rm -f ./sudoku

build:
	go build -ldflags "-X main.buildTime=${shell date --iso-8601=seconds} -X main.gitHash=${shell git rev-parse HEAD}"

test:
	go test ./...
