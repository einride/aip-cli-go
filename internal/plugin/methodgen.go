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

	f.Pf("var %s %s.%s", methodInputVarName(m.method), pbPkg, m.method.Input().Name())

	f.Pf("var %s = &%s.Command{", methodCmdVarName(m.method), cobra)
	f.Pfq("Use: %s,", methodCmdName(m.method))
	f.Pf("RunE: func(cmd *%s.Command, args []string) error {", cobra)
	f.P("conn, err := connect(cmd.Context())")
	f.P("if err != nil {")
	f.P("return err")
	f.P("}")
	f.Pf("client := %s.New%sClient(conn)", pbPkg, m.service.Name())
	f.Pf("result, err := client.%s(cmd.Context(), &%s)", m.method.Name(), methodInputVarName(m.method))
	f.Pf("if err != nil {")
	f.Pf("return err")
	f.Pf("}")
	f.P("_ = result")
	f.Pf("%s.Printf(\"%%v\", result)", fmtPkg)
	f.P("return nil")
	f.P("},")
	f.P("}")
	for k := 0; k < m.method.Input().Fields().Len(); k++ {
		field := m.method.Input().Fields().Get(k)
		if !isSupportedOptionType(field) {
			continue
		}
		f.Pf("var %s %s", optionFlagVarName(m.method, field), optionFlagVarType(field))
	}
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
