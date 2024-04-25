// Package main entrypoint of application
package main

import (
	"github.com/trevatk/chaaya/cmd"
	_ "github.com/trevatk/chaaya/cmd/wallet"
	"go.uber.org/fx"
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
