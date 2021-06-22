package protoflag

import (
	"strconv"
	"strings"

	protoflagv1 "go.einride.tech/protoc-gen-go-cli/proto/gen/go/einride/protoflag/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Int32List struct {
	Value []int32
}

var _ Value = &Int32List{}

func (d *Int32List) String() string {
	var result strings.Builder
	for i, value := range d.Value {
		result.WriteString(strconv.FormatInt(int64(value), 10))
		if i < len(d.Value)-1 {
			result.WriteByte(',')
		}
	}
	return result.String()
}

func (d *Int32List) Set(s string) error {
	ss := strings.Split(s, ",")
	d.Value = make([]int32, 0, len(ss))
	for _, sv := range ss {
		value, err := strconv.ParseInt(sv, 10, 32)
		if err != nil {
			return err
		}
		d.Value = append(d.Value, int32(value))
	}
	return nil
}

func (d *Int32List) Type() string {
	return "[int32]"
}

func (d *Int32List) ProtoReflectValue() protoreflect.Value {
	tmp := protoflagv1.Syntax{RepeatedInt32: d.Value}
	return tmp.ProtoReflect().Get(tmp.ProtoReflect().Descriptor().Fields().ByName("repeated_int32"))
}
