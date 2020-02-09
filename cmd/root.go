package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	buildTime string
	gitHash   string

	rootCmd = &cobra.Command{
		Use:   "sudoku",
		Short: "Solve sudoku puzzles from the command line.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			verbose, _ := cmd.Flags().GetBool("verbose")
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
	}
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetReportCaller(true)

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}

// Execute executes the root command.
func Execute(bTime, gHash string) error {
	buildTime = bTime
	gitHash = gHash
	return rootCmd.Execute()
}
