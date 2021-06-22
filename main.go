package main

import (
	"github.com/spf13/pflag"
	"go.einride.tech/protoc-gen-go-cli/cli"
	"go.einride.tech/protoc-gen-go-cli/internal/gencli"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	flagSet := pflag.NewFlagSet("protoc-gen-go-cli", pflag.ContinueOnError)
	var config cli.CompilerConfig
	config.AddToFlagSet(flagSet)
	protogen.Options{ParamFunc: flagSet.Set}.Run(func(gen *protogen.Plugin) error {
		return gencli.Run(gen, config)
	})
}
