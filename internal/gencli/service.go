package gencli

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type newServiceCommandCodeGenerator struct {
	gen     *protogen.Plugin
	files   *protoregistry.Files
	file    *protogen.File
	service *protogen.Service
}

func (c newServiceCommandCodeGenerator) goName() string {
	return "New" + c.service.GoName + "Command"
}

func (c newServiceCommandCodeGenerator) generateCode(g *protogen.GeneratedFile) error {
	newServiceCommand := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "go.einride.tech/aip-cli/aipcli",
		GoName:       "NewServiceCommand",
	})
	newMethodCommand := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "go.einride.tech/aip-cli/aipcli",
		GoName:       "NewMethodCommand",
	})
	cobraCommand := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "github.com/spf13/cobra",
		GoName:       "Command",
	})
	_ = g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "google.golang.org/protobuf/reflect/protoreflect",
		GoName:       "FullName",
	})
	g.P("func ", c.goName(), "() *", cobraCommand, " {")
	serviceComments := map[protoreflect.FullName]string{}
	collectDescriptorComments(serviceComments, c.service.Desc)
	g.P("cmd := ", newServiceCommand, "(")
	g.P(c.file.GoDescriptorIdent, ".")
	g.P("Services().ByName(\"", c.service.Desc.Name(), "\"),")
	g.P(fmt.Sprintf("%#v", serviceComments), ",")
	g.P(")")
	for _, method := range c.service.Methods {
		methodComments := map[protoreflect.FullName]string{}
		collectMethodComments(methodComments, method.Desc)
		g.P("cmd.AddCommand(")
		g.P(newMethodCommand, "(")
		g.P(c.file.GoDescriptorIdent, ".")
		g.P("Services().ByName(\"", c.service.Desc.Name(), "\").Methods().ByName(\"", method.Desc.Name(), "\"),")
		g.P("&", method.Input.GoIdent, "{},")
		g.P("&", method.Output.GoIdent, "{},")
		g.P(fmt.Sprintf("%#v", methodComments), ",")
		g.P("),")
		g.P(")")
	}
	g.P("return cmd")
	g.P("}")
	return nil
}
