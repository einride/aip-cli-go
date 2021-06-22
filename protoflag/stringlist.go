package protoflag

import (
	"strings"

	protoflagv1 "go.einride.tech/protoc-gen-go-cli/proto/gen/go/einride/protoflag/v1"
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
	tmp := protoflagv1.Syntax{RepeatedString: v.Value}
	return tmp.ProtoReflect().Get(tmp.ProtoReflect().Descriptor().Fields().ByName("repeated_string"))
}
