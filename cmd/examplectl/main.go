package main

import (
	"fmt"
	"os"

	"go.einride.tech/aip-cli/aipcli"
	examplectl "go.einride.tech/aip-cli/cmd/examplectl/gen"
)

func main() {
	if err := examplectl.NewModuleCommand(
		"examplectl",
		"Example CLI tool",
		aipcli.NewIAMModuleCommand("iam", examplectl.NewConfig()),
	).Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
