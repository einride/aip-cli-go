package plugin

import (
	"fmt"
	"strconv"
	"strings"

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
	default:
		return ""
	}
}
