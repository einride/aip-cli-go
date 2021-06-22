package protoflag

import (
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Timestamp struct {
	Value time.Time
}

var _ Value = &Timestamp{}

func (t *Timestamp) String() string {
	return t.Value.UTC().Format(time.RFC3339)
}

func (t *Timestamp) Set(s string) error {
	value, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	t.Value = value.UTC()
	return nil
}

func (t *Timestamp) Type() string {
	return "time"
}

func (t *Timestamp) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfMessage(timestamppb.New(t.Value).ProtoReflect())
}
