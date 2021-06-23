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
		client := g.QualifiedGoIdent(protogen.GoIdent{
			GoImportPath: file.GoImportPath,
			GoName:       service.GoName + "Client",
		})
		parseDialConfig := g.QualifiedGoIdent(protogen.GoIdent{
			GoImportPath: "github.com/einride/ctl",
			GoName:       "ParseDialConfig",
		})
		dial := g.QualifiedGoIdent(protogen.GoIdent{
			GoImportPath: "github.com/einride/ctl",
			GoName:       "Dial",
		})
		newClient := g.QualifiedGoIdent(protogen.GoIdent{
			GoImportPath: file.GoImportPath,
			GoName:       "New" + service.GoName + "Client",
		})
		g.P("// ", service.Desc.FullName(), ".")
		g.P("var (")
		g.P(serviceClientVariableGoName(service), " ", client)
		g.P(serviceCommandVariableGoName(service), " = &", cobraCommand, "{")
		g.P("Use: ", strconv.Quote(string(service.Desc.FullName())), ",")
		g.P("PersistentPreRunE: func(cmd *", cobraCommand, ", args []string) error {")
		g.P("config, err := ", parseDialConfig, "(cmd.Flags())")
		g.P("if err != nil {")
		g.P("return err")
		g.P("}")
		g.P("conn, err := ", dial, "(cmd.Context(), config)")
		g.P("if err != nil {")
		g.P("return err")
		g.P("}")
		g.P(serviceClientVariableGoName(service), " = ", newClient, "(conn)")
		g.P("return nil")
		g.P("},")
		g.P("}")
		g.P(")")
		for _, method := range service.Methods {
			g.P()
			fmtPrintln := g.QualifiedGoIdent(protogen.GoIdent{
				GoImportPath: "fmt",
				GoName:       "Println",
			})
			protojsonFormat := g.QualifiedGoIdent(protogen.GoIdent{
				GoName:       "Format",
				GoImportPath: "google.golang.org/protobuf/encoding/protojson",
			})
			g.P("// ", method.Desc.FullName(), ".")
			g.P("var (")
			g.P(requestVariableGoName(method), " ", method.Input.GoIdent)
			g.P(methodCommandVariableGoName(method), " = &", cobraCommand, "{")
			g.P("Use: ", strconv.Quote(string(method.Desc.Name())), ",")
			g.P("RunE: func(cmd *", cobraCommand, ", args []string) error {")
			g.P("response, err := ", serviceClientVariableGoName(service), ".", method.GoName, "(cmd.Context(), &", requestVariableGoName(method), ")")
			g.P("if err != nil {")
			g.P("return err")
			g.P("}")
			g.P(fmtPrintln, "(", protojsonFormat, "(response))")
			g.P("return nil")
			g.P("},")
			g.P("}")
			g.P(")")
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
			for _, field := range method.Input.Fields {
				e := env{
					g:      g,
					method: method,
					field:  field,
				}
				g.P()
				e.generate()
			}
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
	g.P("flags := Command.PersistentFlags()")
	g.P(`flags.Bool("insecure", false, "make insecure client connection, must be used with 'address' option")`)
	g.P(`flags.Bool("prod", false, "connect to prod")`)
	g.P(`flags.String("address", "", "address to connect to")`)
	g.P(`flags.String("token", "", "bearer token used by client")`)
	g.P(`flags.String("project", "einride-dev", "GCP project")`)
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
	g.P("AddCompletion(Command)")
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

func serviceClientVariableGoName(service *protogen.Service) string {
	return strings.ReplaceAll(string(service.Desc.FullName()), ".", "_") + "Client"
}

func serviceConfigVariableGoName(service *protogen.Service) string {
	return strings.ReplaceAll(string(service.Desc.FullName()), ".", "_") + "Config"
}

func serviceConfigInitFunctionGoName(service *protogen.Service) string {
	return "init_" + strings.ReplaceAll(string(service.Desc.FullName()), ".", "_") + "Config"
}

func methodCommandVariableGoName(method *protogen.Method) string {
	return strings.ReplaceAll(string(method.Desc.FullName()), ".", "_")
}

func requestVariableGoName(method *protogen.Method) string {
	return strings.ReplaceAll(string(method.Desc.FullName()), ".", "_") + "_Request"
}

func addCommandFunctionGoName(service *protogen.Service) string {
	return "Add" + string(service.GoName) + "Command"
}
