package messagebroker

import (
	"github.com/spf13/cobra"
	"github.com/trevatk/chaaya/cmd"
)

var (
	msgCmd = &cobra.Command{
		Use: "message",
	}
)

func init() {
	cmd.RootCmd.AddCommand(msgCmd)
}
