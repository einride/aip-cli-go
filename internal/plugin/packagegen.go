package plugin

import (
	"strings"

	"github.com/einride/ctl/internal/codegen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type PackageGenerator struct {
	name  protoreflect.FullName
	files []protoreflect.FileDescriptor
}

func (p PackageGenerator) Generate(f *codegen.File) error {
	f.P("package ctl")
	f.P(`import (
	"fmt"
	"github.com/spf13/cobra"
)`)
	f.P("var _ = fmt.Sprintf")
	f.P("var _ = cobra.Command{}")
	for _, file := range p.files {
		for i := 0; i < file.Services().Len(); i++ {
			service := file.Services().Get(i)
			f.Pf("var %s = &cobra.Command{", serviceCmdVarName(service))
			f.Pfq("Use: %s,", serviceCmdName(service))
			f.P("Run: func(cmd *cobra.Command, args []string) {")
			f.Pfq("fmt.Println(%s)", serviceCmdName(service)+" called")
			f.P("},")
			f.P("}")
			f.P()

			for j := 0; j < service.Methods().Len(); j++ {
				method := service.Methods().Get(j)
				f.Pf("var %s = &cobra.Command{", methodCmdVarName(method))
				f.Pfq("Use: %s,", methodCmdName(method))
				f.P("RunE: func(cmd *cobra.Command, args []string) error {")
				f.Pfq("fmt.Println(%s)", methodCmdName(method)+" called")
				f.P("conn, err := connect()")
				f.P("if err != nil {")
				f.P("return err")
				f.P("}")
				f.P("client := accountv1beta1.NewFeatureFlagsServiceClient(conn)")
				f.P("_ = client")
				f.P("return nil")
				f.P("},")
				f.P("}")
				f.P()
			}
		}
	}

	f.P("func init() {")
	for _, file := range p.files {
		for i := 0; i < file.Services().Len(); i++ {
			service := file.Services().Get(i)
			f.Pf("rootCmd.AddCommand(%s)", serviceCmdVarName(service))
			for j := 0; j < service.Methods().Len(); j++ {
				method := service.Methods().Get(j)
				f.Pf("%s.AddCommand(%s)", serviceCmdVarName(service), methodCmdVarName(method))
			}
		}
	}
	f.P("}")

	return nil
}

func serviceCmdVarName(service protoreflect.ServiceDescriptor) string {
	segments := strings.Split(string(service.FullName()), ".")
	return strings.Join(segments, "_")
}

func serviceCmdName(service protoreflect.ServiceDescriptor) string {
	segments := strings.Split(string(service.FullName()), ".")
	segments[len(segments)-1] = strings.Replace(segments[len(segments)-1], "Service", "", 1)
	return strings.Join(segments, ".")
}

func methodCmdVarName(method protoreflect.MethodDescriptor) string {
	segments := strings.Split(string(method.FullName()), ".")
	return strings.Join(segments, "_")
}

func methodCmdName(method protoreflect.MethodDescriptor) string {
	return string(method.Name())
}
