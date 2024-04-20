package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "chaaya",
		Short: "Chaaya command line client for structx blockchain",
		Long:  "",
	}
)

func Execute() error {
	return RootCmd.Execute()
}
