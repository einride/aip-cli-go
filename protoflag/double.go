package protoflag

import (
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Double struct {
	Value float64
}

var _ Value = &Double{}

func (d *Double) String() string {
	return strconv.FormatFloat(d.Value, 'f', -1, 64)
}

func (d *Double) Set(s string) error {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	d.Value = value
	return nil
}

func (d *Double) Type() string {
	return "float64"
}

func (d *Double) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfFloat64(d.Value)
}
