package cmd

import (
	"github.com/spf13/cobra"
)

var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "operations involving EC2 instances",
	Long:  "non-destructive operations involving EC2 instances",
}

func init() {
	rootCmd.AddCommand(instanceCmd)
}
