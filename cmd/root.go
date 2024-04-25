// Package cmd application commands
package cmd

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var (
	// RootCmd cobra cli root command
	RootCmd = &cobra.Command{
		Use:   "chaaya",
		Short: "Chaaya command line client for structx blockchain",
		Long:  "",
	}
)

// Action uberfx cli action
type Action struct {
	sh fx.Shutdowner
}

// NewAction constructor
func NewAction(lc fx.Lifecycle, sh fx.Shutdowner) *Action {
	a := &Action{
		sh: sh,
	}
	lc.Append(fx.Hook{
		OnStart: a.start,
		OnStop:  a.stop,
	})
	return a
}

func (a *Action) start(ctx context.Context) error {
	go a.run(ctx)
	return nil
}

func (a *Action) stop(_ context.Context) error {
	return nil
}

// execute cobra root command
func (a *Action) run(ctx context.Context) {
	log.SetTimeFormat(time.Kitchen)
	log.SetLevel(log.DebugLevel)

	timeout, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	if err := RootCmd.ExecuteContext(timeout); err != nil {
		log.Error("failed to execute command", "error", err)
	}
	_ = a.sh.Shutdown()
}
