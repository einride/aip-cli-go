package plugin

import (
	"fmt"
	"strconv"

	"github.com/stoewer/go-strcase"

	"github.com/einride/ctl/internal/codegen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MethodGenerator struct {
	service protoreflect.ServiceDescriptor
	method  protoreflect.MethodDescriptor
}

func (m MethodGenerator) GenerateCmd(f *codegen.File) {
	cobra := f.Import("github.com/spf13/cobra")
	fmtPkg := f.Import("fmt")
	pbPkg := f.Import(getGoPkg(m.method).path)
	protoJSON := f.Import("google.golang.org/protobuf/encoding/protojson")

	f.Pf("var %s %s.%s", methodInputVarName(m.method), pbPkg, m.method.Input().Name())

	f.P("var ", methodCmdVarName(m.method), " = &", cobra, ".Command{")
	f.P("Use: ", strconv.Quote(methodCmdName(m.method)), ",")
	f.P("RunE: func(cmd *", cobra, ".Command, args []string) error {")
	f.P("conn, err := connect(cmd.Context())")
	f.P("if err != nil {")
	f.P("return err")
	f.P("}")
	f.P("client := ", pbPkg, ".New", m.service.Name(), "Client(conn)")
	f.Pf("result, err := client.%s(cmd.Context(), &%s)", m.method.Name(), methodInputVarName(m.method))
	f.Pf("if err != nil {")
	f.Pf("return err")
	f.Pf("}")
	f.P(fmtPkg, ".Println(", protoJSON, ".Format(result))")
	f.P("return nil")
	f.P("},")
	f.P("}")
}

func (m MethodGenerator) GenerateInit(f *codegen.File) {
	f.Pf("%s.AddCommand(%s)", serviceCmdVarName(m.service), methodCmdVarName(m.method))
	for k := 0; k < m.method.Input().Fields().Len(); k++ {
		field := m.method.Input().Fields().Get(k)
		if !isSupportedOptionType(field) {
			continue
		}
		f.Pf(
			"%s.Flags().%s(&%s, %s, %s, %s)",
			methodCmdVarName(m.method),
			optionFlagCreateFunc(field),
			fmt.Sprintf("%s.%s", methodInputVarName(m.method), strcase.UpperCamelCase(string(field.Name()))),
			strconv.Quote(optionFlagName(field)),
			optionDefaultValue(field),
			strconv.Quote(optionFlagDescription(field)),
		)
	}
	f.P()
}
