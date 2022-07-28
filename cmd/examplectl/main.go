package main

import (
	"fmt"
	"os"

	examplectl "go.einride.tech/aip-cli/cmd/examplectl/gen"
)

func main() {
	if err := examplectl.NewModuleCommand().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
