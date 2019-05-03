package main

import (
	"os"

	"github.com/FracKenA/op5beat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
