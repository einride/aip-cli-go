package plugin

import (
	"github.com/einride/ctl/internal/codegen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type PackageGenerator struct {
	name  protoreflect.FullName
	files []protoreflect.FileDescriptor
}

func (p PackageGenerator) Generate(f *codegen.File) error {
	if len(p.files) == 0 {
		return nil
	}
	goPkg := getGoPkg(p.files[0])

	f.Pf(`import (
	"fmt"
	"github.com/spf13/cobra"
	%s "%s"
)`, goPkg.name, goPkg.path)
	f.P("var _ = fmt.Sprintf")
	f.P("var _ = cobra.Command{}")
	f.P("var _ = cobra.Command{}")

	for _, file := range p.files {
		for i := 0; i < file.Services().Len(); i++ {
			service := file.Services().Get(i)
			ServiceGenerator{service: service}.GenerateCmd(f)
			f.P()
		}
	}

	f.P("func init() {")
	for _, file := range p.files {
		for i := 0; i < file.Services().Len(); i++ {
			service := file.Services().Get(i)
			ServiceGenerator{service: service}.GenerateInit(f)
		}
	}
	f.P("}")

	return nil
}
