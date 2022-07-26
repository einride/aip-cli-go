package gencli

import (
	"fmt"
	"path"
	"strings"

	"go.einride.tech/aip-cli/aipcli"
	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateMainFile(gen *protogen.Plugin, config aipcli.CompilerConfig) error {
	module, ok := getModuleParam(gen)
	if !ok {
		return fmt.Errorf("param main requires param module to be provided")
	}
	g := gen.NewGeneratedFile(path.Join(module, "main.go"), "")
	generateGeneratedFileHeader(g, gen)
	cliWithConfig := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "go.einride.tech/aip-cli/aipcli",
		GoName:       "WithConfig",
	})
	contextBackground := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "context",
		GoName:       "Background",
	})
	g.P("package main")
	g.P()
	g.P("func main() {")
	g.P("ctx := ", contextBackground, "()")
	g.P("cmd := NewRootCommand()")
	g.P("config := NewConfig()")
	g.P("config.Runtime.AddToFlagSet(cmd.PersistentFlags(), config.Compiler)")
	g.P("ctx = ", cliWithConfig, "(ctx, config)")
	g.P("if err := cmd.ExecuteContext(ctx); err != nil {")
	g.P(
		"_, _ = ",
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "fmt", GoName: "Fprintln"}),
		"(",
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "os", GoName: "Stderr"}),
		", err)",
	)
	g.P(g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "os", GoName: "Exit"}), "(1)")
	g.P("}")
	g.P("}")
	return nil
}

func getModuleParam(gen *protogen.Plugin) (string, bool) {
	for _, param := range strings.Split(gen.Request.GetParameter(), ",") {
		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = param[i+1:]
			param = param[0:i]
		}
		if param == "module" {
			return value, true
		}
	}
	return "", false
}
