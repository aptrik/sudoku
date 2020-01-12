package main

import (
	"sudoku/cmd"
)

var (
	buildTime = "?"
	gitHash   = "?"
)

func main() {
	cmd.Execute(buildTime, gitHash)
}
