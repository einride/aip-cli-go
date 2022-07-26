package main

import (
	"github.com/spf13/pflag"
	"go.einride.tech/aip-cli/aipcli"
	"go.einride.tech/aip-cli/internal/gencli"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	flagSet := pflag.NewFlagSet("aip-cli", pflag.ContinueOnError)
	var config aipcli.CompilerConfig
	config.AddToFlagSet(flagSet)
	protogen.Options{ParamFunc: flagSet.Set}.Run(func(gen *protogen.Plugin) error {
		return gencli.Run(gen, config)
	})
}
