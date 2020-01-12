package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sudoku/solver"
)

func init() {
	rootCmd.AddCommand(solveCmd)
}

var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		solveF := func(filename string)  {
			f, _ := os.Open(filename)
			//noinspection GoUnhandledErrorResult
			defer f.Close()
			reader := bufio.NewReader(f)
			board, err := solver.NewBoardFrom(reader)
			if err != nil {
				fmt.Printf("Unable to create board from input: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("Problem %q:\n", filename)
			fmt.Println(board)
			solution, _ := board.Solve()
			fmt.Println()
			fmt.Printf("Solution to %q:\n", filename)
			fmt.Println(solution)
			fmt.Println()
		}
		for _, filename := range args {
			solveF(filename)
		}
	},
}
