package genctl

import (
	"log"
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const generateFilenameSuffix = "_einridectl.pb.go"

func GenerateFile(gen *protogen.Plugin, file *protogen.File) error {
	g := gen.NewGeneratedFile(
		file.GeneratedFilenamePrefix+generateFilenameSuffix,
		// TODO: Remove this patch when generating from example code not in a descriptor.
		protogen.GoImportPath(strings.ReplaceAll(
			string(file.GoImportPath),
			"github.com/einride/proto/gen/go",
			"github.com/einride/ctl/example/ctl2",
		)),
	)
	g.P("package ", file.GoPackageName)
	g.Skip()
	cobraCommand := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "github.com/spf13/cobra",
		GoName:       "Command",
	})
	for _, service := range file.Services {
		g.Unskip()
		g.P()
		g.P("var ", serviceCommandVariableGoName(service), " = &", cobraCommand, "{")
		g.P("Use: ", strconv.Quote(string(service.Desc.FullName())), ",")
		g.P("}")
		for _, method := range service.Methods {
			g.P()
			request := g.QualifiedGoIdent(protogen.GoIdent{
				GoImportPath: file.GoImportPath,
				GoName:       method.Input.GoIdent.GoName,
			})
			logPrintln := g.QualifiedGoIdent(protogen.GoIdent{
				GoImportPath: "log",
				GoName:       "Println",
			})
			g.P("var ", requestVariableGoName(method), " ", request)
			g.P("var ", methodCommandVariableGoName(method), " = &", cobraCommand, "{")
			g.P("Use: ", strconv.Quote(string(method.Desc.Name())), ",")
			g.P("RunE: func(cmd *", cobraCommand, ", args []string) error {")
			g.P(logPrintln, "(", strconv.Quote(string(method.Desc.FullName())), ")")
			g.P("return nil")
			g.P("},")
			g.P("}")
		}
	}
	for _, service := range file.Services {
		g.P()
		g.P("func ", addCommandFunctionGoName(service), "(parent *", cobraCommand, ") {")
		g.P("parent.AddCommand(", serviceCommandVariableGoName(service), ")")
		g.P("}")
	}
	g.P()
	g.P("func init() {")
	for _, service := range file.Services {
		for _, method := range service.Methods {
			g.P(serviceCommandVariableGoName(service), ".AddCommand(", methodCommandVariableGoName(method), ")")
		}
	}
	g.P("}")
	return nil
}

func GenerateRootFile(gen *protogen.Plugin, rootPackage string) error {
	module := getModule(gen)
	log.Println(rootPackage)
	log.Println(module)
	g := gen.NewGeneratedFile(module+"/root.go", "")
	cobraCommand := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "github.com/spf13/cobra",
		GoName:       "Command",
	})
	g.P("package ", rootPackage)
	g.P()
	g.P("var Command = &", cobraCommand, "{")
	g.P("Use: ", strconv.Quote("einridectl2"), ",")
	g.P("}")
	g.P()
	g.P("func init() {")
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		for _, service := range file.Services {
			addServiceCommandFunction := g.QualifiedGoIdent(protogen.GoIdent{
				GoName: addCommandFunctionGoName(service),
				GoImportPath: protogen.GoImportPath(strings.ReplaceAll(
					string(file.GoImportPath),
					"github.com/einride/proto/gen/go",
					"github.com/einride/ctl/example/ctl2",
				)),
			})
			g.P(addServiceCommandFunction, "(Command)")
		}
	}
	g.P("}")
	return nil
}

func getModule(gen *protogen.Plugin) string {
	for _, param := range strings.Split(gen.Request.GetParameter(), ",") {
		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = param[i+1:]
			param = param[0:i]
		}
		if param == "module" {
			return value
		}
	}
	return ""
}

func serviceCommandVariableGoName(service *protogen.Service) string {
	return strings.ReplaceAll(string(service.Desc.FullName()), ".", "_")
}

func methodCommandVariableGoName(method *protogen.Method) string {
	return strings.ReplaceAll(string(method.Desc.FullName()), ".", "_")
}

func requestVariableGoName(method *protogen.Method) string {
	return strings.ReplaceAll(string(method.Input.Desc.FullName()), ".", "_")
}

func addCommandFunctionGoName(service *protogen.Service) string {
	return "Add" + string(service.GoName) + "Command"
}
