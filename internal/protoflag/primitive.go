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
	return &primitiveValue[T]{
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
	strVal  string
}

// String returns the last value passed to Set, or "" if Set has not been called.
// This makes cmd.Flags().GetString() work correctly for proto string fields.
func (v *primitiveValue[T]) String() string {
	return v.strVal
}

func (v *primitiveValue[T]) Set(s string) error {
	parsed, err := v.parser(s)
	if err != nil {
		return err
	}
	v.mutable().Set(v.field, v.valueOf(parsed))
	v.strVal = s
	return nil
}

func (v *primitiveValue[T]) Type() string {
	return v.field.Kind().String()
}
