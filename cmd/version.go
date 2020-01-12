package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Golang Env: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Build time: %s\n", buildTime)
		fmt.Printf("Build hash: %s\n", gitHash)
	},
}
