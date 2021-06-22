package protoflag

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type String struct {
	Value string
}

var _ Value = &String{}

func (v *String) String() string {
	return v.Value
}

func (v *String) Set(s string) error {
	v.Value = s
	return nil
}

func (v *String) Type() string {
	return "string"
}

func (v *String) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfString(v.String())
}
