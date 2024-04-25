// Package wallet command line operations
package wallet

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	wallet "github.com/trevatk/go-wallet"
)

var (
	outputFile string

	createWalletCmd = &cobra.Command{
		Use:   "create",
		Short: "create wallet",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Helper()
			log.Info("execute command", "cmd", cmd.Use, "args", args, "flags", []any{outputFile})

			w := wallet.New()
			err := w.MarshalToFile(outputFile)
			if err != nil {
				return fmt.Errorf("failed to marshal wallet %v", err)
			}

			return nil
		},
	}
)

func init() {

	createWalletCmd.Flags().StringVarP(&outputFile, "output-file", "o", "wallet.json", "set output file location")

	walletCmd.AddCommand(createWalletCmd)
}
