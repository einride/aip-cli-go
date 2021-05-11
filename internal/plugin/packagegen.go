package plugin

import (
	"github.com/einride/ctl/internal/codegen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
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
			f.P("var ", serviceCmdVarName(service), " = &cobra.Command{")
			f.P("Use: \"", serviceCmdName(service), "\",")
			f.P("Run: func(cmd *cobra.Command, args []string) {")
			f.P("fmt.Println(\"", serviceCmdName(service)," called\")")
			f.P("},")
			f.P("}")
		}
	}

	f.P("func init() {")
	for _, file := range p.files {
		for i := 0; i < file.Services().Len(); i++ {
			service := file.Services().Get(i)
			f.P("rootCmd.AddCommand(", serviceCmdVarName(service),")")
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
