package wallet_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/structx/cmdline-client/cmd"
)

func Test_ExecuteWallet(t *testing.T) {
	t.Run("help", func(t *testing.T) {

		b := bytes.NewBuffer(nil)
		cmd.RootCmd.SetOutput(b)
		cmd.RootCmd.SetArgs([]string{"wallet"})

		err := cmd.Execute()
		if err != nil {
			t.Fatalf("failed to execute command")
		}

		content, err := io.ReadAll(b)
		if err != nil {
			t.Fatalf("unable to read output %v", err)
		}

		if bytes.Equal([]byte{}, content) {
			t.Fatal("no output")
		}
	})
}
