package main

import (
	"github.com/aptrik/sudoku/cmd"
)

var (
	buildTime = "?"
	gitHash   = "?"
)

func main() {
	cmd.Execute(buildTime, gitHash)
}
