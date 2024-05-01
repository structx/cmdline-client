// Package main entrypoint of application
package main

import (
	"github.com/charmbracelet/log"

	"github.com/structx/cmdline-client/cmd"
	_ "github.com/structx/cmdline-client/cmd/wallet"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Errorf("failed to execute command %v", err)
	}
}
