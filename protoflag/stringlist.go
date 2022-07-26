package protoflag

import (
	"strings"

	aipcliv1 "go.buf.build/protocolbuffers/go/einride/aip-cli/einride/aipcli/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type StringList struct {
	Value []string
}

var _ Value = &StringList{}

func (v *StringList) String() string {
	return strings.Join(v.Value, ",")
}

func (v *StringList) Set(s string) error {
	v.Value = strings.Split(s, ",")
	return nil
}

func (v *StringList) Type() string {
	return "[string]"
}

func (v *StringList) ProtoReflectValue() protoreflect.Value {
	tmp := aipcliv1.Values{RepeatedString: v.Value}
	return tmp.ProtoReflect().Get(tmp.ProtoReflect().Descriptor().Fields().ByName("repeated_string"))
}
