package plugin

import (
	"strconv"

	"github.com/einride/ctl/internal/codegen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MethodGenerator struct {
	service protoreflect.ServiceDescriptor
	method  protoreflect.MethodDescriptor
}

func (m MethodGenerator) GenerateCmd(f *codegen.File) {
	f.Pf("var %s = &cobra.Command{", methodCmdVarName(m.method))
	f.Pfq("Use: %s,", methodCmdName(m.method))
	f.P("RunE: func(cmd *cobra.Command, args []string) error {")
	f.Pfq("fmt.Println(%s)", methodCmdName(m.method)+" called")
	f.P("conn, err := connect(cmd.Context())")
	f.P("if err != nil {")
	f.P("return err")
	f.P("}")
	f.P("_ = conn")
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
			optionFlagVarName(m.method, field),
			strconv.Quote(optionFlagName(field)),
			optionDefaultValue(field),
			strconv.Quote(optionFlagDescription(field)),
		)
	}
}
