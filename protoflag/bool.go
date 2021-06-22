package protoflag

import (
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Bool struct {
	Value bool
}

var _ Value = &Bool{}

func (v *Bool) String() string {
	return strconv.FormatBool(v.Value)
}

func (v *Bool) Set(s string) error {
	value, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	v.Value = value
	return nil
}

func (v *Bool) Type() string {
	return "bool"
}

func (v *Bool) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfBool(v.Value)
}
