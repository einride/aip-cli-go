package protoflag

import (
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Value interface {
	pflag.Value

	// ProtoReflectValue returns a protoreflect.Value representation of the flag value.
	ProtoReflectValue() protoreflect.Value
}
