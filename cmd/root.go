// Package cmd application commands
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// RootCmd cobra cli root command
	RootCmd = &cobra.Command{
		Use:   "chaaya",
		Short: "Chaaya command line client for structx blockchain",
		Long:  "",
	}
)

// Execute root command
func Execute() error {
	return RootCmd.Execute()
}
