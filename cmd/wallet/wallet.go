package wallet

import (
	"github.com/spf13/cobra"
	"github.com/trevatk/chaaya/cmd"
)

var (
	walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "create wallet",
	}
)

func init() {
	cmd.RootCmd.AddCommand(walletCmd)
}
