package main

import (
	"fmt"
	"os"

	"go.einride.tech/aip-cli/aipcli"
	deployctl "go.einride.tech/aip-cli/cmd/deployctl/gen"
)

func main() {
	// It is also possible to pass options to NewModule(), for deploy
	//
	// deployctl.NewModule(
	// 	aipcli.WithGlobalTokenFunc(
	// 		func(cmd *cobra.Command) (string, error) {
	//			return <my bearer token>, nil
	// 		},
	// 	),
	// )
	if err := deployctl.NewModule().Command(
		"deployctl",
		"deploy CLI tool",
		aipcli.NewIAMModuleCommand("iam", deployctl.NewConfig()),
	).Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
