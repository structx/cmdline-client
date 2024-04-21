package wallet_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/teatest"
	"github.com/trevatk/chaaya/cmd"
	"github.com/trevatk/chaaya/cmd/wallet"
)

func Test_ExecuteCreateWallet(t *testing.T) {
	t.Run("create wallet default output", func(t *testing.T) {

		cmd.RootCmd.SetArgs([]string{"wallet", "create"})

		m := wallet.CreateWalletModel{}
		tm := teatest.NewTestModel(t, m)

		_ = cmd.Execute()

		teatest.WaitFor(
			t,
			tm.Output(),
			func(bts []byte) bool {
				return true
			},
		)
	})
}
