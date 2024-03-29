// Code generated by protoc-gen-go-aip-cli. DO NOT EDIT.
package examplectl

import (
	cobra "github.com/spf13/cobra"
	aipcli "go.einride.tech/aip-cli/aipcli"
	v1 "go.einride.tech/aip-cli/cmd/examplectl/gen/einride/example/freight/v1"
)

func NewModuleCommand(use string, short string, commands ...*cobra.Command) *cobra.Command {
	config := NewConfig()
	return aipcli.NewModuleCommand(
		use,
		short,
		config,
		append(
			[]*cobra.Command{
				v1.NewFreightServiceCommand(config),
			},
			commands...,
		)...,
	)
}

func NewConfig() aipcli.Config {
	return aipcli.Config{Hosts: map[string]string{}, DefaultHost: "", RootPath: "", Root: "examplectl", GoogleCloudIdentityTokens: true, CachedIdentityTokenPath: ""}
}
