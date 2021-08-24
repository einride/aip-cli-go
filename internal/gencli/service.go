package gencli

import (
	"strconv"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
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
	cobraCommand := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "github.com/spf13/cobra",
		GoName:       "Command",
	})
	cliSetDefaultHost := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "go.einride.tech/protoc-gen-go-cli/cli",
		GoName:       "SetDefaultHost",
	})
	g.P()
	g.P("func ", c.goName(), "(use string) *", cobraCommand, " {")
	g.P("cmd := &", cobraCommand, "{")
	g.P("Use: use,")
	g.P("Short: ", strconv.Quote(getShortService(c.service)), ",")
	g.P("Annotations: map[string]string{")
	g.P(strconv.Quote("type"), ": ", strconv.Quote("service"), ",")
	g.P("},")
	if defaultHost := proto.GetExtension(
		c.service.Desc.Options(), annotations.E_DefaultHost,
	).(string); defaultHost != "" {
		g.P("PersistentPreRun: func(cmd *", cobraCommand, ", _ []string) {")
		g.P(cliSetDefaultHost, "(cmd.Context(), ", strconv.Quote(defaultHost), ")")
		g.P("},")
	}
	g.P("}")
	for _, method := range c.service.Methods {
		g.P("cmd.AddCommand(", c.newMethodCommand(method).goName(), "())")
	}
	g.P("return cmd")
	g.P("}")
	for _, method := range c.service.Methods {
		if err := (c.newMethodCommand(method)).generateCode(g); err != nil {
			return err
		}
	}
	return nil
}

func (c newServiceCommandCodeGenerator) newMethodCommand(method *protogen.Method) newMethodCommandCodeGenerator {
	return newMethodCommandCodeGenerator{
		gen:     c.gen,
		file:    c.file,
		files:   c.files,
		service: c.service,
		method:  method,
	}
}
