package main

import (
	"github.com/spf13/pflag"
	"go.einride.tech/aip-cli/internal/gengoaipcli"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	flagSet := pflag.NewFlagSet("aip-cli", pflag.ContinueOnError)
	var config gengoaipcli.Config
	config.AddToFlagSet(flagSet)
	protogen.Options{ParamFunc: flagSet.Set}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return gengoaipcli.Run(gen, config)
	})
}
