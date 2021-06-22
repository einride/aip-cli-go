package protoflag

import (
	"encoding/base64"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Bytes struct {
	Value []byte
}

var _ Value = &Bytes{}

func (b *Bytes) String() string {
	return base64.StdEncoding.EncodeToString(b.Value)
}

func (b *Bytes) Set(s string) error {
	value, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	b.Value = value
	return nil
}

func (b *Bytes) Type() string {
	return "base64"
}

func (b *Bytes) ProtoReflectValue() protoreflect.Value {
	return protoreflect.ValueOfBytes(b.Value)
}
