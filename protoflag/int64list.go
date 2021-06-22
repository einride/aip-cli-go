package protoflag

import (
	"strconv"
	"strings"

	protoflagv1 "go.einride.tech/protoc-gen-go-cli/proto/gen/go/einride/protoflag/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Int64List struct {
	Value []int64
}

var _ Value = &Int64List{}

func (d *Int64List) String() string {
	var result strings.Builder
	for i, value := range d.Value {
		result.WriteString(strconv.FormatInt(value, 10))
		if i < len(d.Value)-1 {
			result.WriteByte(',')
		}
	}
	return result.String()
}

func (d *Int64List) Set(s string) error {
	ss := strings.Split(s, ",")
	d.Value = make([]int64, 0, len(ss))
	for _, sv := range ss {
		value, err := strconv.ParseInt(sv, 10, 64)
		if err != nil {
			return err
		}
		d.Value = append(d.Value, value)
	}
	return nil
}

func (d *Int64List) Type() string {
	return "[int64]"
}

func (d *Int64List) ProtoReflectValue() protoreflect.Value {
	tmp := protoflagv1.Syntax{RepeatedInt64: d.Value}
	return tmp.ProtoReflect().Get(tmp.ProtoReflect().Descriptor().Fields().ByName("repeated_int64"))
}
