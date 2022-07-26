package protoflag

import (
	"strconv"
	"strings"

	aipcliv1 "go.buf.build/protocolbuffers/go/einride/aip-cli/einride/aipcli/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type BoolList struct {
	Value []bool
}

var _ Value = &BoolList{}

func (v *BoolList) String() string {
	var result strings.Builder
	for i, value := range v.Value {
		result.WriteString(strconv.FormatBool(value))
		if i < len(v.Value)-1 {
			result.WriteByte(',')
		}
	}
	return result.String()
}

func (v *BoolList) Set(s string) error {
	ss := strings.Split(s, ",")
	v.Value = make([]bool, 0, len(ss))
	for _, sv := range ss {
		value, err := strconv.ParseBool(sv)
		if err != nil {
			return err
		}
		v.Value = append(v.Value, value)
	}
	return nil
}

func (v *BoolList) Type() string {
	return "[bool]"
}

func (v *BoolList) ProtoReflectValue() protoreflect.Value {
	tmp := aipcliv1.Values{RepeatedBool: v.Value}
	return tmp.ProtoReflect().Get(tmp.ProtoReflect().Descriptor().Fields().ByName("repeated_bool"))
}
