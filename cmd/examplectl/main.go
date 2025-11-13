package main

import (
	"fmt"
	"os"

	"go.einride.tech/aip-cli/aipcli"
	examplectl "go.einride.tech/aip-cli/cmd/examplectl/gen"
)

func main() {
	// It is also possible to pass options to NewModule(), for example
	//
	// examplectl.NewModule(
	// 	aipcli.WithGlobalTokenFunc(
	// 		func(cmd *cobra.Command) (string, error) {
	//			return <my bearer token>, nil
	// 		},
	// 	),
	// )
	if err := examplectl.NewModule().Command(
		"examplectl",
		"Example CLI tool",
		aipcli.NewIAMModuleCommand("iam", examplectl.NewConfig()),
	).Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
