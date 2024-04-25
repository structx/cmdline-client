package messagebroker

import "github.com/spf13/cobra"

var (
	pingCmd = &cobra.Command{
		Use:  "ping",
		RunE: ping,
	}
)

func init() {
	msgCmd.AddCommand(pingCmd)
}

func ping(cmd *cobra.Command, args []string) error {
	return nil
}
