package main

import (
	"flag"

	"github.com/einride/ctl/internal/genctl"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	flagSet := flag.NewFlagSet("protoc-gen-einridectl", flag.ExitOnError)
	var rootPackage string
	flagSet.StringVar(&rootPackage, "rootPackage", "", "TODO")
	protogen.Options{
		ParamFunc: flagSet.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				if err := genctl.GenerateFile(gen, f); err != nil {
					return err
				}
			}
		}
		if err := genctl.GenerateRootFile(gen, rootPackage); err != nil {
			return err
		}
		return nil
	})
}
