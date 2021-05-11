package plugin

import (
	"fmt"
	"strings"

	"github.com/einride/ctl/internal/codegen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func Generate(request *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	generate := make(map[string]struct{})
	registry, err := protodesc.NewFiles(&descriptorpb.FileDescriptorSet{
		File: request.ProtoFile,
	})
	if err != nil {
		return nil, fmt.Errorf("create proto registry: %w", err)
	}
	for _, f := range request.FileToGenerate {
		generate[f] = struct{}{}
	}
	packaged := make(map[protoreflect.FullName][]protoreflect.FileDescriptor)
	for _, f := range request.FileToGenerate {
		file, err := registry.FindFileByPath(f)
		if err != nil {
			return nil, fmt.Errorf("find file %s: %w", f, err)
		}
		packaged[file.Package()] = append(packaged[file.Package()], file)
	}

	var res pluginpb.CodeGeneratorResponse
	for pkg, files := range packaged {
		var index codegen.File
		pkgElements := strings.Split(string(pkg), ".")
		filename := strings.Join(pkgElements, "_")

		if err := (PackageGenerator{
			name:  pkg,
			files: files,
		}.Generate(&index)); err != nil {
			return nil, err
		}
		content, err := index.Content()
		if err != nil {
			return nil, fmt.Errorf("generated content: %w", err)
		}
		res.File = append(res.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    proto.String(filename + ".go"),
			Content: proto.String(string(content)),
		})
	}

	var rootFile codegen.File
	RootGenerator{}.Generate(&rootFile)
	content, err := rootFile.Content()
	if err != nil {
		return nil, fmt.Errorf("generated content: %w", err)
	}
	res.File = append(res.File, &pluginpb.CodeGeneratorResponse_File{
		Name:    proto.String("root.go"),
		Content: proto.String(string(content)),
	})

	return &res, nil
}
