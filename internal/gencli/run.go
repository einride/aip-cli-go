package gencli

import (
	"go.einride.tech/protoc-gen-go-cli/cli"
	"google.golang.org/protobuf/compiler/protogen"
)

type Config struct {
	RootCommand string
	Hosts       map[string]string
	DefaultHost string
}

func Run(gen *protogen.Plugin, config cli.CompilerConfig) error {
	if err := config.Validate(); err != nil {
		return err
	}
	for _, f := range gen.Files {
		if f.Generate {
			if err := GenerateFile(gen, f, config); err != nil {
				return err
			}
		}
	}
	if config.Root != "" {
		if err := GenerateRootCommandFile(gen, config); err != nil {
			return err
		}
		if config.Main {
			if err := GenerateMainFile(gen, config); err != nil {
				return err
			}
		}
	}
	return nil
}
