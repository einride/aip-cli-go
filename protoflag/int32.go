package protoflag

import (
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Int32 struct {
	Value int32
}

var _ Value = &Int32{}

func (i *Int32) String() string {
	return strconv.FormatInt(int64(i.Value), 10)
}

func (i *Int32) Set(s string) error {
	value, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	i.Value = int32(value)
	return nil
}

func (i *Int32) Type() string {
	return "int32"
}

func (i *Int32) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfInt32(i.Value)
}
