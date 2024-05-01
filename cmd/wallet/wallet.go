package wallet

import (
	"github.com/spf13/cobra"
	"github.com/structx/cmdline-client/cmd"
)

var (
	walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "wallet subcommands",
		Long:  "All available wallet subcommands supported within the application",
	}
)

func init() {
	cmd.RootCmd.AddCommand(walletCmd)
}
