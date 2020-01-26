package main

import (
	"github.com/aptrik/sudoku/cmd"
	"os"
)

var (
	buildTime = "?"
	gitHash   = "?"
)

func main() {
	if err := cmd.Execute(buildTime, gitHash); err != nil {
		os.Exit(1)
	}
}
