# Sudoku Solver

## Install

    go get -u github.com/aptrik/sudoku

## Build

    make build

## Test

    make test

## Update Go dependencies

    go get -u ./...
    go mod tidy

## Usage

`./sudoku solve puzzle1.txt`

    *** Problem "puzzle1.txt":
    0 0 5 0 0 0 4 0 9
    0 0 0 6 0 1 3 0 8
    0 3 0 2 0 0 1 0 5
    5 0 4 0 1 0 6 0 0
    2 7 0 0 0 0 0 4 0
    1 8 0 0 4 7 0 9 3
    7 0 0 0 3 9 0 1 6
    9 6 0 1 2 5 7 0 0
    0 0 0 0 6 8 9 5 0

    +++ Solution in 52.924µs:
    6 1 5 8 7 3 4 2 9
    4 2 9 6 5 1 3 7 8
    8 3 7 2 9 4 1 6 5
    5 9 4 3 1 2 6 8 7
    2 7 3 9 8 6 5 4 1
    1 8 6 5 4 7 2 9 3
    7 5 2 4 3 9 8 1 6
    9 6 8 1 2 5 7 3 4
    3 4 1 7 6 8 9 5 2
