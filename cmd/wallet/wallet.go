package wallet

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/trevatk/chaaya/cmd"
)

var (
	walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "wallet",
		Long:  "All available wallet subcommands supported within the application",
		Run: func(cmd *cobra.Command, _ []string) {
			log.Helper()
			log.Info("execute command", "cmd", cmd.Use, "args", cmd.Args)
		},
	}
)

func init() {
	cmd.RootCmd.AddCommand(walletCmd)
}
