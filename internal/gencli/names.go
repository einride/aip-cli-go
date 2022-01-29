package gencli

import (
	"strings"

	"github.com/stoewer/go-strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
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

func getMethodCommandUse(method *protogen.Method) string {
	return strcase.KebabCase(string(method.Desc.Name()))
}

func getServiceCommandUse(servicesByName map[protoreflect.Name][]*protogen.Service, service *protogen.Service) string {
	var serviceName string
	if services := servicesByName[service.Desc.Name()]; len(services) > 1 {
		serviceName = strings.TrimSuffix(string(service.Desc.FullName()), "Service")
		serviceName = strings.TrimPrefix(serviceName, string(longestCommonParent(services)))
		serviceName = strings.TrimPrefix(serviceName, ".")
		serviceName = strings.Join(reverse(strings.Split(serviceName, ".")), "-")
	} else {
		serviceName = strings.TrimSuffix(string(service.Desc.Name()), "Service")
	}
	return strcase.KebabCase(serviceName)
}

func getServicesByName(gen *protogen.Plugin) map[protoreflect.Name][]*protogen.Service {
	result := map[protoreflect.Name][]*protogen.Service{}
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		for _, service := range file.Services {
			result[service.Desc.Name()] = append(result[service.Desc.Name()], service)
		}
	}
	return result
}

// longestCommonParent returns the longest common parent of multiple services.
func longestCommonParent(services []*protogen.Service) protoreflect.FullName {
	if len(services) == 0 {
		return ""
	}
	var i int
	var result protoreflect.FullName
ResultLoop:
	for {
		parts := strings.Split(string(services[0].Desc.FullName()), ".")
		if i > len(parts) {
			break
		}
		candidate := strings.Join(parts[:i], ".")
		for _, service := range services {
			if !strings.HasPrefix(string(service.Desc.FullName()), candidate) {
				break ResultLoop
			}
		}
		result = protoreflect.FullName(candidate)
		i++
	}
	return result
}

func reverse(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}
