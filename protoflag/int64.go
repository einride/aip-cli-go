package protoflag

import (
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Int64 struct {
	Value int64
}

var _ Value = &Int64{}

func (i *Int64) String() string {
	return strconv.FormatInt(i.Value, 10)
}

func (i *Int64) Set(s string) error {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	i.Value = value
	return nil
}

func (i *Int64) Type() string {
	return "int64"
}

func (i *Int64) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfInt64(i.Value)
}
