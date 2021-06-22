package protoflag

import (
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Duration struct {
	Value time.Duration
}

func (d *Duration) String() string {
	return d.Value.String()
}

func (d *Duration) Set(s string) error {
	value, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.Value = value
	return nil
}

func (d *Duration) Type() string {
	return "duration"
}

func (d *Duration) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfMessage(durationpb.New(d.Value).ProtoReflect())
}
