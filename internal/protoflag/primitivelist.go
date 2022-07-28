package protovalue

import (
	"strings"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func PrimitiveList[T any](
	mutable func() protoreflect.Message,
	field protoreflect.FieldDescriptor,
	valueOf func(T) protoreflect.Value,
	parser func(string) (T, error),
) pflag.Value {
	return primitiveListValue[T]{
		mutable: mutable,
		field:   field,
		parser:  parser,
		valueOf: valueOf,
	}
}

type primitiveListValue[T any] struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
	parser  func(string) (T, error)
	valueOf func(T) protoreflect.Value
}

func (v primitiveListValue[T]) String() string {
	return ""
}

func (v primitiveListValue[T]) Set(s string) error {
	values := strings.Split(s, ",")
	if len(values) == 0 {
		return nil
	}
	msg := v.mutable()
	list := msg.NewField(v.field).List()
	for _, value := range values {
		parsed, err := v.parser(value)
		if err != nil {
			return err
		}
		list.Append(v.valueOf(parsed))
	}
	msg.Set(v.field, protoreflect.ValueOfList(list))
	return nil
}

func (v primitiveListValue[T]) Type() string {
	return "[" + v.field.Kind().String() + "]"
}
