package plugin

import (
	"github.com/einride/ctl/internal/codegen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func Generate(request *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	var response pluginpb.CodeGeneratorResponse

	generate := make(map[string]struct{})
	for _, f := range request.FileToGenerate {
		var file codegen.File
		generate[f] = struct{}{}
		file.P("// ", f)
		response.File = append(response.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    proto.String(f + ".go"),
			Content: proto.String(string(file.Content())),
		})
	}

	return &response, nil
}
