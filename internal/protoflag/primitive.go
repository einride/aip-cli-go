package protovalue

import (
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Primitive[T any](
	mutable func() protoreflect.Message,
	field protoreflect.FieldDescriptor,
	valueOf func(T) protoreflect.Value,
	parser func(string) (T, error),
) pflag.Value {
	return primitiveValue[T]{
		mutable: mutable,
		field:   field,
		valueOf: valueOf,
		parser:  parser,
	}
}

type primitiveValue[T any] struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
	valueOf func(T) protoreflect.Value
	parser  func(string) (T, error)
}

func (v primitiveValue[T]) String() string {
	return ""
}

func (v primitiveValue[T]) Set(s string) error {
	parsed, err := v.parser(s)
	if err != nil {
		return err
	}
	v.mutable().Set(v.field, v.valueOf(parsed))
	return nil
}

func (v primitiveValue[T]) Type() string {
	return v.field.Kind().String()
}
