// Package main entrypoint of application
package main

import (
	"go.uber.org/fx"

	"github.com/trevatk/chaaya/cmd"
	_ "github.com/trevatk/chaaya/cmd/messagebroker"
	_ "github.com/trevatk/chaaya/cmd/wallet"
)

func main() {
	fx.New(
		fx.Provide(
			cmd.NewAction,
		),
		fx.Invoke(func(*cmd.Action) {}),
		fx.NopLogger,
	).Run()
}
