package protovalue

import (
	"fmt"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Enum(mutable func() protoreflect.Message, field protoreflect.FieldDescriptor) pflag.Value {
	return enumValue{mutable: mutable, field: field}
}

type enumValue struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
}

func (v enumValue) String() string {
	return ""
}

func (v enumValue) Type() string {
	return "enum[" + string(v.field.Enum().Name()) + "]"
}

func (v enumValue) Set(s string) error {
	value := v.field.Enum().Values().ByName(protoreflect.Name(s))
	if value == nil {
		return fmt.Errorf("no such value for %v: %v", v.field.Enum().Name(), s)
	}
	v.mutable().Set(v.field, protoreflect.ValueOfEnum(value.Number()))
	return nil
}
