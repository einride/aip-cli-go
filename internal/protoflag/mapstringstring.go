package protovalue

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func MapStringString(mutable func() protoreflect.Message, field protoreflect.FieldDescriptor) pflag.Value {
	return mapStringStringValue{mutable: mutable, field: field}
}

type mapStringStringValue struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
}

func (v mapStringStringValue) String() string {
	return ""
}

func (v mapStringStringValue) Type() string {
	return "map<string, string>"
}

func (v mapStringStringValue) Set(s string) error {
	pairs := strings.Split(s, ",")
	if len(pairs) == 0 {
		return nil
	}
	value := v.mutable().NewField(v.field)
	for _, pair := range pairs {
		keyValue := strings.SplitN(pair, "=", 2)
		if len(keyValue) != 2 {
			return fmt.Errorf("invalid map pair: %s", pair)
		}
		value.Map().Set(protoreflect.ValueOfString(keyValue[0]).MapKey(), protoreflect.ValueOfString(keyValue[1]))
	}
	v.mutable().Set(v.field, value)
	return nil
}
