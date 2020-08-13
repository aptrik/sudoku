package main

import (
	"os"

	"github.com/aptrik/sudoku/cmd"
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
