package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/aptrik/sudoku/solver"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(solveCmd)
}

var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		solveF := func(filename string) {
			f, _ := os.Open(filename)
			//noinspection GoUnhandledErrorResult
			defer f.Close()
			reader := bufio.NewReader(f)
			board, err := solver.NewBoardFrom(reader)
			if err != nil {
				fmt.Printf("Unable to create board from input: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("*** Problem %q:\n", filename)
			fmt.Print(board)
			fmt.Println()
			start := time.Now()
			solution, err := board.Solve()
			elapsed := time.Since(start)
			if err != nil {
				fmt.Printf("!!! Error: %v", err)
				fmt.Println()
			} else {
				fmt.Printf("+++ Solution in %v:\n", elapsed)
				fmt.Print(solution)
				fmt.Println()
			}
		}
		for _, filename := range args {
			solveF(filename)
		}
	},
}
