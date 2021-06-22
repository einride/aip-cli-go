package gencli

import (
	"strings"

	"github.com/stoewer/go-strcase"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

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
		i += 1
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
