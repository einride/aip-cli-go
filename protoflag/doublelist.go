package protoflag

import (
	"strconv"
	"strings"

	protoflagv1 "go.einride.tech/protoc-gen-go-cli/proto/gen/go/einride/protoflag/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DoubleList struct {
	Value []float64
}

var _ Value = &DoubleList{}

func (d *DoubleList) String() string {
	var result strings.Builder
	for i, value := range d.Value {
		result.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
		if i < len(d.Value)-1 {
			result.WriteByte(',')
		}
	}
	return result.String()
}

func (d *DoubleList) Set(s string) error {
	ss := strings.Split(s, ",")
	d.Value = make([]float64, 0, len(ss))
	for _, sv := range ss {
		value, err := strconv.ParseFloat(sv, 64)
		if err != nil {
			return err
		}
		d.Value = append(d.Value, value)
	}
	return nil
}

func (d *DoubleList) Type() string {
	return "[float64]"
}

func (d *DoubleList) ProtoReflectValue() protoreflect.Value {
	tmp := protoflagv1.Syntax{RepeatedDouble: d.Value}
	return tmp.ProtoReflect().Get(tmp.ProtoReflect().Descriptor().Fields().ByName("repeated_double"))
}
