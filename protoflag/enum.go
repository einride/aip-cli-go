package protoflag

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Enum struct {
	Value protoreflect.Enum
}

func (e *Enum) String() string {
	return string(e.Value.Descriptor().Values().ByNumber(e.Value.Number()).Name())
}

func (e *Enum) Set(s string) error {
	value := e.Value.Descriptor().Values().ByName(protoreflect.Name(s))
	if value == nil {
		return fmt.Errorf("no such value for %v: %v", e.Value.Type().Descriptor().Name(), s)
	}
	e.Value = e.Value.Type().New(value.Number())
	return nil
}

func (e *Enum) Type() string {
	return fmt.Sprintf("enum[%s]", e.Value.Type().Descriptor().Name())
}

func (e *Enum) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfEnum(e.Value.Number())
}
