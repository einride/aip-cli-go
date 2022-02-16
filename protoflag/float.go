package protoflag

import (
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Float struct {
	Value float32
}

var _ Value = &Float{}

func (d *Float) String() string {
	return strconv.FormatFloat(float64(d.Value), 'f', -1, 32)
}

func (d *Float) Set(s string) error {
	value, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	d.Value = float32(value)
	return nil
}

func (d *Float) Type() string {
	return "float32"
}

func (d *Float) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfFloat32(d.Value)
}
