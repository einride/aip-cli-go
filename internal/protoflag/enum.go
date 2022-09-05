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

func lookupEnum(field protoreflect.FieldDescriptor, name string) (protoreflect.EnumValueDescriptor, error) {
	value := field.Enum().Values().ByName(protoreflect.Name(name))
	if value == nil {
		return nil, fmt.Errorf("no such value for %v: %v", field.Enum().Name(), name)
	}
	return value, nil
}

func (v enumValue) Set(s string) error {
	value, err := lookupEnum(v.field, s)
	if err != nil {
		return err
	}

	v.mutable().Set(v.field, protoreflect.ValueOfEnum(value.Number()))
	return nil
}

func EnumList(mutable func() protoreflect.Message, field protoreflect.FieldDescriptor) pflag.Value {
	parser := func(s string) (protoreflect.EnumValueDescriptor, error) {
		return lookupEnum(field, s)
	}
	valueOf := func(v protoreflect.EnumValueDescriptor) protoreflect.Value {
		return protoreflect.ValueOfEnum(v.Number())
	}
	return PrimitiveList[protoreflect.EnumValueDescriptor](mutable, field, valueOf, parser)
}
