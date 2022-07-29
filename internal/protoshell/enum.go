package protoshell

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// CompleteEnumValue returns valid values to complete for the provided enum values.
func CompleteEnumValue(toComplete string, values protoreflect.EnumValueDescriptors) []string {
	result := make([]string, 0, values.Len())
	for i := 0; i < values.Len(); i++ {
		value := string(values.Get(i).Name())
		if strings.HasPrefix(value, toComplete) {
			result = append(result, value)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
