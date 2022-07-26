package gencli

import (
	"go.einride.tech/aip-cli/aipcli"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type Config struct {
	RootCommand string
	Hosts       map[string]string
	DefaultHost string
}

func Run(gen *protogen.Plugin, config aipcli.CompilerConfig) error {
	var files protoregistry.Files
	for _, file := range gen.Files {
		if err := files.RegisterFile(file.Desc); err != nil {
			return err
		}
	}
	if err := config.Validate(); err != nil {
		return err
	}
	for _, f := range gen.Files {
		if f.Generate {
			if err := GenerateFile(gen, &files, f, config); err != nil {
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
