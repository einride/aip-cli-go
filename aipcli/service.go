package aipcli

import (
	"strings"

	"github.com/stoewer/go-strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func getServiceCommandUse(
	servicesByName map[protoreflect.Name][]protoreflect.FullName,
	service protoreflect.FullName,
) string {
	var serviceName string
	if services := servicesByName[service.Name()]; len(services) > 1 {
		serviceName = strings.TrimSuffix(string(service), "Service")
		serviceName = strings.TrimPrefix(serviceName, string(longestCommonParent(services)))
		serviceName = strings.TrimPrefix(serviceName, ".")
		serviceName = strings.Join(reverse(strings.Split(serviceName, ".")), "-")
	} else {
		serviceName = strings.TrimSuffix(string(service.Name()), "Service")
	}
	return strcase.KebabCase(serviceName)
}

// longestCommonParent returns the longest common parent of multiple services.
func longestCommonParent(fullNames []protoreflect.FullName) protoreflect.FullName {
	if len(fullNames) == 0 {
		return ""
	}
	var i int
	var result protoreflect.FullName
ResultLoop:
	for {
		parts := strings.Split(string(fullNames[0]), ".")
		if i > len(parts) {
			break
		}
		candidate := strings.Join(parts[:i], ".")
		for _, fullName := range fullNames {
			if !strings.HasPrefix(string(fullName), candidate) {
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
