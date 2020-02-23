all: clean build test

LDFLAGS := -X main.buildTime=${shell date --iso-8601=seconds} -X main.gitHash=${shell git rev-parse HEAD}

clean:
	rm -f ./sudoku

build:
	go build -ldflags "${LDFLAGS}"

test:
	go test -v ./...

bench:
	go test -v -bench=. -benchmem -run=^bench ./...
