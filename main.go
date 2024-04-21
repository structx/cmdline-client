// Package main entrypoint of application
package main

import (
	"github.com/charmbracelet/log"

	"github.com/trevatk/chaaya/cmd"
	_ "github.com/trevatk/chaaya/cmd/wallet"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Errorf("failed to execute command %v", err)
	}
}
