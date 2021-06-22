package protoflag

import (
	"fmt"
	"strings"

	protoflagv1 "go.einride.tech/protoc-gen-go-cli/proto/gen/go/einride/protoflag/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type StringStringMap struct {
	Value map[string]string
}

var _ Value = &StringStringMap{}

func (m *StringStringMap) String() string {
	pairs := make([]string, 0, len(m.Value))
	for k, v := range m.Value {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(pairs, ",")
}

func (m *StringStringMap) Set(s string) error {
	pairs := strings.Split(s, ",")
	m.Value = make(map[string]string, len(pairs))
	for _, pair := range pairs {
		keyValue := strings.SplitN(pair, "=", 2)
		if len(keyValue) != 2 {
			return fmt.Errorf("invalid map pair: %s", pair)
		}
		m.Value[keyValue[0]] = keyValue[1]
	}
	return nil
}

func (m *StringStringMap) Type() string {
	return "map<string,string>"
}

func (m *StringStringMap) ProtoReflectValue() protoreflect.Value {
	tmp := protoflagv1.Syntax{MapStringString: m.Value}
	return tmp.ProtoReflect().Get(tmp.ProtoReflect().Descriptor().Fields().ByName("map_string_string"))
}
