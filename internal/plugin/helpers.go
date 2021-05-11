package plugin

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

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

func optionFlagVarName(method protoreflect.MethodDescriptor, field protoreflect.FieldDescriptor) string {
	return fmt.Sprintf("%s_%s", methodCmdVarName(method), field.Name())
}

func optionFlagName(field protoreflect.FieldDescriptor) string {
	return string(field.Name())
}

func optionFlagDescription(field protoreflect.FieldDescriptor) string {
	return field.ParentFile().SourceLocations().ByDescriptor(field).LeadingComments
}

func isSupportedOptionType(field protoreflect.FieldDescriptor) bool {
	if field.IsList() || field.IsMap() {
		return false
	}
	switch field.Kind() {
	case protoreflect.StringKind:
		return true
	case protoreflect.Int32Kind:
		return true
	case protoreflect.BoolKind:
		return true
	default:
		return false
	}
}

func optionFlagVarType(field protoreflect.FieldDescriptor) string {
	switch field.Kind() {
	case protoreflect.StringKind:
		return "string"
	case protoreflect.Int32Kind:
		return "int32"
	case protoreflect.BoolKind:
		return "bool"
	default:
		return ""
	}
}

func optionFlagCreateFunc(field protoreflect.FieldDescriptor) string {
	switch field.Kind() {
	case protoreflect.StringKind:
		return "StringVar"
	case protoreflect.Int32Kind:
		return "Int32Var"
	case protoreflect.BoolKind:
		return "BoolVar"
	default:
		return ""
	}
}

func optionDefaultValue(field protoreflect.FieldDescriptor) string {
	switch field.Kind() {
	case protoreflect.StringKind:
		return strconv.Quote("")
	case protoreflect.Int32Kind:
		return "0"
	case protoreflect.BoolKind:
		return "false"
	default:
		return ""
	}
}

type goPkg struct {
	path string
	name string
}

func getGoPkg(desc protoreflect.Descriptor) goPkg {
	pkg := protodesc.ToFileDescriptorProto(desc.ParentFile()).Options.GetGoPackage()

	if strings.Contains(pkg, ";") {
		parts := strings.Split(pkg, ";")
		return goPkg{
			path: parts[0],
			name: parts[1],
		}
	}
	return goPkg{
		path: pkg,
		name: "pb",
	}
}
