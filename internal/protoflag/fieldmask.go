package protovalue

import (
	"strings"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func FieldMask(mutable func() protoreflect.Message, field protoreflect.FieldDescriptor) pflag.Value {
	return fieldMaskValue{mutable: mutable, field: field}
}

type fieldMaskValue struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
}

func (v fieldMaskValue) String() string {
	return ""
}

func (v fieldMaskValue) Type() string {
	return "field-mask"
}

func (v fieldMaskValue) Set(s string) error {
	values := strings.Split(s, ",")
	if len(values) == 0 {
		return nil
	}
	fieldMask := fieldmaskpb.FieldMask{
		Paths: make([]string, 0, len(values)),
	}
	for _, value := range values {
		fieldMask.Paths = append(fieldMask.Paths, strings.TrimSpace(value))
	}
	v.mutable().Set(v.field, protoreflect.ValueOf(fieldMask.ProtoReflect()))
	return nil
}
