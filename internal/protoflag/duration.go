package protovalue

import (
	"time"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
)

func Duration(mutable func() protoreflect.Message, field protoreflect.FieldDescriptor) pflag.Value {
	return durationValue{mutable: mutable, field: field}
}

type durationValue struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
}

func (v durationValue) String() string {
	return ""
}

func (v durationValue) Type() string {
	return "duration"
}

func (v durationValue) Set(s string) error {
	duration, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	v.mutable().Set(v.field, protoreflect.ValueOf(durationpb.New(duration).ProtoReflect()))
	return nil
}
