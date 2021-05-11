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
	f.P("// ", p.name)
	for _, file := range p.files {
		f.P("// ", file.Path())
	}
	return nil
}
