package wallet_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/structx/cmdline-client/cmd"
	"github.com/structx/cmdline-client/cmd/wallet"
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
				return !bytes.Equal([]byte{}, bts)
			},
		)

		tm.Send(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune("q"),
		})
	})
}
