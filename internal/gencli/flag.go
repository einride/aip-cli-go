package gencli

import (
	"strings"

	"github.com/stoewer/go-strcase"
	"google.golang.org/protobuf/compiler/protogen"
)

func getFlagName(field *protogen.Field, parents []*protogen.Field) string {
	var result strings.Builder
	for _, parent := range parents {
		_, _ = result.WriteString(strcase.KebabCase(string(parent.Desc.Name())))
		_ = result.WriteByte('.')
	}
	_, _ = result.WriteString(strcase.KebabCase(string(field.Desc.Name())))
	return result.String()
}

func getFlagVariableName(field *protogen.Field, parents []*protogen.Field) string {
	var result strings.Builder
	_, _ = result.WriteString("flag_")
	for _, parent := range parents {
		_, _ = result.WriteString(parent.GoName)
		_ = result.WriteByte('_')
	}
	_, _ = result.WriteString(field.GoName)
	return result.String()
}

func getFlagDescription(field *protogen.Field) string {
	return trimComments(field.Comments.Leading)
}

func getShortService(service *protogen.Service) string {
	return trimComments(service.Comments.Leading)
}

func getShortMethod(method *protogen.Method) string {
	return trimComments(method.Comments.Leading)
}

func trimComments(comments protogen.Comments) string {
	result := comments.String()
	// Clean up comment line breaks and whitespace.
	result = strings.ReplaceAll(result, "//", "")
	result = strings.ReplaceAll(result, "\n", " ")
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "  ", " ")
	result = strings.ReplaceAll(result, "  ", " ")
	// Cut out first sentence.
	if i := strings.IndexByte(result, '.'); i != -1 {
		result = result[:i]
	}
	// Trim manually documented field behavior.
	result = strings.TrimPrefix(result, "REQUIRED: ")
	result = strings.TrimPrefix(result, "Required: ")
	// Lower-case.
	result = strings.ToLower(result)
	return result
}
