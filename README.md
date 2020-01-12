# Sudoku Solver

## Install

    go get -u github.com/aptrik/sudoku

## Release

    go build -ldflags "-X main.buildTime=$(date --iso-8601=seconds) -X main.gitHash=$(git rev-parse HEAD)"

## Update Go dependencies

    go get -u ./...
    go mod tidy

## Test

    go test ./...

## Usage

cat puzzle1.txt
```
  0 2 3 6 5 0 0 0 4
  0 1 4 2 3 0 8 0 0
  9 7 5 8 0 0 0 0 0
  4 0 0 7 0 0 0 0 0
  5 3 2 0 0 0 9 4 7
  7 0 8 0 2 5 0 3 0
  1 4 7 0 8 0 0 5 0
  0 0 0 0 1 0 0 8 0
  3 8 9 0 0 0 0 1 2
```

    sudoku solve puzzle1.txt
```
+-------+-------+-------+
| 6 1 5 | 8 7 3 | 4 2 9 |
| 4 2 9 | 6 5 1 | 3 7 8 |
| 8 3 7 | 2 9 4 | 1 6 5 |
+-------+-------+-------+
| 5 9 4 | 3 1 2 | 6 8 7 |
| 2 7 3 | 9 8 6 | 5 4 1 |
| 1 8 6 | 5 4 7 | 2 9 3 |
+-------+-------+-------+
| 7 5 2 | 4 3 9 | 8 1 6 |
| 9 6 8 | 1 2 5 | 7 3 4 |
| 3 4 1 | 7 6 8 | 9 5 2 |
+-------+-------+-------+
```
