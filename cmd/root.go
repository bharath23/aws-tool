package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "awstool",
	Short:        "awstool - a simple CLI to perform non-destructive operations on aws",
	Long:         "awstool is a simple tool to query and perform non-destructive operations on aws",
	SilenceUsage: true,
	Version:      "0.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
