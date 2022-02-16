package protoflag

import (
	"strconv"
	"strings"

	protoflagv1 "go.einride.tech/protoc-gen-go-cli/proto/gen/go/einride/protoflag/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FloatList struct {
	Value []float32
}

var _ Value = &FloatList{}

func (d *FloatList) String() string {
	var result strings.Builder
	for i, value := range d.Value {
		result.WriteString(strconv.FormatFloat(float64(value), 'f', -1, 64))
		if i < len(d.Value)-1 {
			result.WriteByte(',')
		}
	}
	return result.String()
}

func (d *FloatList) Set(s string) error {
	ss := strings.Split(s, ",")
	d.Value = make([]float32, 0, len(ss))
	for _, sv := range ss {
		value, err := strconv.ParseFloat(sv, 64)
		if err != nil {
			return err
		}
		d.Value = append(d.Value, float32(value))
	}
	return nil
}

func (d *FloatList) Type() string {
	return "[float32]"
}

func (d *FloatList) ProtoReflectValue() protoreflect.Value {
	tmp := protoflagv1.Syntax{RepeatedFloat: d.Value}
	return tmp.ProtoReflect().Get(tmp.ProtoReflect().Descriptor().Fields().ByName("repeated_float"))
}
