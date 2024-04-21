// Package main entrypoint of application
package main

import (
	"os"

	"github.com/trevatk/chaaya/cmd"
	_ "github.com/trevatk/chaaya/cmd/wallet"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
