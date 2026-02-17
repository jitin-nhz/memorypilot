package main

import (
	"os"

	"github.com/memorypilot/memorypilot/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
