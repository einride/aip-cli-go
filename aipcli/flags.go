package aipcli

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	protovalue "go.einride.tech/aip-cli/internal/protoflag"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const (
	FieldResourceName = "name"
)

func setFlags(
	comments map[protoreflect.FullName]string,
	cmd *cobra.Command,
	parentFields []protoreflect.FieldDescriptor,
	msg protoreflect.MessageDescriptor,
	mutable func() protoreflect.Message,
	requiredFlagNames *[]string,
) {
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		if isOutputOnly(field) {
			// Field is a OUTPUT_ONLY field, omit generating any flags.
			continue
		}
		switch field.Kind() {
		case protoreflect.MessageKind:
			switch field.Message().FullName() {
			case "google.protobuf.Duration":
				if !field.IsList() {
					addFlag(cmd, field, parentFields, comments[field.FullName()], protovalue.Duration(mutable, field), requiredFlagNames)
				}
				// TODO: Implement support for repeated durations.
			case "google.protobuf.Timestamp":
				if !field.IsList() {
					addFlag(cmd, field, parentFields, comments[field.FullName()], protovalue.Timestamp(mutable, field), requiredFlagNames)
				}
				// TODO: Implement support for repeated timestamps.
			case "google.protobuf.FieldMask":
				if !field.IsList() {
					addFlag(cmd, field, parentFields, comments[field.FullName()], protovalue.FieldMask(mutable, field), requiredFlagNames)
				}
				// Repeated field masks is intentionally not supported.
			default:
				switch {
				case field.IsMap():
					switch {
					case field.MapKey().Kind() == protoreflect.StringKind &&
						field.MapValue().Kind() == protoreflect.StringKind:
						addFlag(
							cmd,
							field,
							parentFields,
							comments[field.FullName()],
							protovalue.MapStringString(mutable, field),
							requiredFlagNames,
						)
					default:
						// TODO: Implement support for more map types.
					}
				case field.IsList():
					// Repeated nested messages not supported.
				default:
					setFlags(
						comments,
						cmd,
						append(parentFields, field),
						field.Message(),
						func() protoreflect.Message {
							return mutable().Mutable(field).Message()
						},
						requiredFlagNames,
					)
				}
			}
		case protoreflect.EnumKind:
			if field.IsList() {
				addFlag(cmd, field, parentFields, comments[field.FullName()], protovalue.EnumList(mutable, field), requiredFlagNames)
			} else {
				addFlag(cmd, field, parentFields, comments[field.FullName()], protovalue.Enum(mutable, field), requiredFlagNames)
			}
		case protoreflect.StringKind, protoreflect.BoolKind, protoreflect.BytesKind, protoreflect.DoubleKind,
			protoreflect.FloatKind, protoreflect.Int64Kind, protoreflect.Int32Kind:
			setPrimitiveFlag(comments, cmd, parentFields, mutable, field, requiredFlagNames)
		}
	}
}

func setPrimitiveFlag(
	comments map[protoreflect.FullName]string,
	cmd *cobra.Command,
	parentFields []protoreflect.FieldDescriptor,
	mutable func() protoreflect.Message,
	field protoreflect.FieldDescriptor,
	requiredFlagNames *[]string,
) {
	var value pflag.Value
	switch field.Kind() {
	case protoreflect.BoolKind:
		if field.IsList() {
			value = protovalue.PrimitiveList[bool](mutable, field, protoreflect.ValueOfBool, strconv.ParseBool)
		} else {
			value = protovalue.Primitive[bool](mutable, field, protoreflect.ValueOfBool, strconv.ParseBool)
		}
	case protoreflect.StringKind:
		parser := func(s string) (string, error) {
			return s, nil
		}
		if field.IsList() {
			value = protovalue.PrimitiveList[string](mutable, field, protoreflect.ValueOfString, parser)
		} else {
			value = protovalue.Primitive[string](mutable, field, protoreflect.ValueOfString, parser)
		}
	case protoreflect.BytesKind:
		value = protovalue.Primitive[[]byte](mutable, field, protoreflect.ValueOfBytes, base64.URLEncoding.DecodeString)
	case protoreflect.DoubleKind:
		parser := func(s string) (float64, error) {
			return strconv.ParseFloat(s, 64)
		}
		if field.IsList() {
			value = protovalue.PrimitiveList[float64](mutable, field, protoreflect.ValueOfFloat64, parser)
		} else {
			value = protovalue.Primitive[float64](mutable, field, protoreflect.ValueOfFloat64, parser)
		}
	case protoreflect.FloatKind:
		parser := func(s string) (float32, error) {
			d, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return 0, err
			}
			return float32(d), nil
		}
		if field.IsList() {
			value = protovalue.PrimitiveList[float32](mutable, field, protoreflect.ValueOfFloat32, parser)
		} else {
			value = protovalue.Primitive[float32](mutable, field, protoreflect.ValueOfFloat32, parser)
		}
	case protoreflect.Int64Kind:
		parser := func(s string) (int64, error) {
			return strconv.ParseInt(s, 10, 64)
		}
		if field.IsList() {
			value = protovalue.PrimitiveList[int64](mutable, field, protoreflect.ValueOfInt64, parser)
		} else {
			value = protovalue.Primitive[int64](mutable, field, protoreflect.ValueOfInt64, parser)
		}
	case protoreflect.Int32Kind:
		parser := func(s string) (int32, error) {
			i64, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return 0, err
			}
			return int32(i64), nil
		}
		if field.IsList() {
			value = protovalue.PrimitiveList[int32](mutable, field, protoreflect.ValueOfInt32, parser)
		} else {
			value = protovalue.Primitive[int32](mutable, field, protoreflect.ValueOfInt32, parser)
		}
	default:
		panic(fmt.Errorf("unhandled primitive kind: %v", field.Kind())) // shouldn't happen
	}
	addFlag(cmd, field, parentFields, comments[field.FullName()], value, requiredFlagNames)
}

func addFlag(
	cmd *cobra.Command,
	field protoreflect.FieldDescriptor,
	parentFields []protoreflect.FieldDescriptor,
	comment string,
	value pflag.Value,
	requiredFlagNames *[]string,
) {
	flag := &pflag.Flag{
		Name:  flagName(field, parentFields),
		Usage: trimFieldComment(comment),
		Value: value,
		Annotations: map[string][]string{
			fieldNameAnnotation: {string(field.FullName())},
		},
	}
	cmd.Flags().AddFlag(flag)
	_ = cmd.Flags().SetAnnotation(flag.Name, flagArgumentAnnotation, []string{})
	annotateFlagWithFieldBehaviors(flag, field, parentFields)
	markRequiredFlags(requiredFlagNames, cmd, flag, field, parentFields)
	hideOutputOnlyFields(cmd, flag, field)
	registerCompletion(cmd, flag, field, comment)
	hideImmutableForUpdateMethods(cmd, flag, field)
	hideNameForCreateMethods(cmd, flag, field)
	hideETagForCreateMethods(cmd, flag, field)
}

func markRequiredFlags(
	requiredFlagNames *[]string,
	cmd *cobra.Command,
	flag *pflag.Flag,
	field protoreflect.FieldDescriptor,
	parentFields []protoreflect.FieldDescriptor,
) {
	if isMethodType(cmd, "Update") {
		// Update methods have no required fields due to field masks.
		return
	}
	if os.Getenv("AIP_CLI_DISABLE_FIELD_BEHAVIOR") == "true" {
		// A secret escape hatch for ignoring required fields.
		return
	}
	if anyParentIsOptional(parentFields) {
		// Omit marking flag as required if any parent is optional, it has higher precedence.
		return
	}
	if fieldBehaviors, ok := proto.GetExtension(
		field.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior); ok {
		for _, fieldBehavior := range fieldBehaviors {
			if fieldBehavior == annotations.FieldBehavior_REQUIRED {
				// Defer the MarkFlagsOneRequired call until after all proto-field flags are
				// registered. Calling it here would trigger mergePersistentFlags(), which
				// copies persistent flags (e.g. --address) into local Flags() before the
				// colliding proto-field flag is added, causing a "flag redefined" panic.
				*requiredFlagNames = append(*requiredFlagNames, flag.Name)
				return
			}
		}
	}
}

func hideOutputOnlyFields(
	cmd *cobra.Command,
	flag *pflag.Flag,
	field protoreflect.FieldDescriptor,
) {
	if isMethodType(cmd, "Update") && field.Name() == FieldResourceName {
		// Update methods need to be able to specify the resource name to update
		return
	}
	if fieldBehaviors, ok := proto.GetExtension(
		field.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior); ok {
		for _, fieldBehavior := range fieldBehaviors {
			if fieldBehavior == annotations.FieldBehavior_OUTPUT_ONLY {
				_ = cmd.Flags().MarkHidden(flag.Name)
				return
			}
		}
	}
}

func annotateFlagWithFieldBehaviors(
	flag *pflag.Flag,
	field protoreflect.FieldDescriptor,
	parentFields []protoreflect.FieldDescriptor,
) {
	if fieldBehaviors, ok := proto.GetExtension(
		field.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior); ok {
		for _, fieldBehavior := range fieldBehaviors {
			if flag.Annotations == nil {
				flag.Annotations = map[string][]string{}
			}
			fb := fieldBehavior
			if fieldBehavior == annotations.FieldBehavior_REQUIRED && anyParentIsOptional(parentFields) {
				// Omit marking flag as required if any parent is optional, it has higher precedence.
				// This to relax the requirements if two fields are optional, but their children has
				// required field behavior.
				fb = annotations.FieldBehavior_OPTIONAL
			}
			flag.Annotations[fieldBehaviorsAnnotation] = append(
				flag.Annotations[fieldBehaviorsAnnotation],
				fb.String(),
			)
		}
	}
}

func registerCompletion(
	cmd *cobra.Command,
	flag *pflag.Flag,
	field protoreflect.FieldDescriptor,
	comment string,
) {
	// resource name fields
	if !field.IsList() && field.Name() == FieldResourceName && field.Kind() == protoreflect.StringKind {
		if resourceDescriptor, ok := proto.GetExtension(
			field.Parent().Options(),
			annotations.E_Resource,
		).(*annotations.ResourceDescriptor); ok && resourceDescriptor.GetType() != "" {
			var didRegisterCompletion bool
			aipreflect.RangeResourceDescriptorsInPackage(
				protoregistry.GlobalFiles,
				field.ParentFile().Package(),
				func(resource *annotations.ResourceDescriptor) bool {
					if resource.GetType() == resourceDescriptor.GetType() && len(resource.GetPattern()) > 0 {
						didRegisterCompletion = true
						_ = cmd.RegisterFlagCompletionFunc(
							flag.Name,
							resourceNameCompletionFunc(comment, resource.GetPattern()...),
						)
						return false
					}
					return true
				},
			)
			if didRegisterCompletion {
				return
			}
		}
	}
	// resource reference fields
	if field.Kind() == protoreflect.StringKind {
		if resourceReference, ok := proto.GetExtension(
			field.Options(),
			annotations.E_ResourceReference,
		).(*annotations.ResourceReference); ok && resourceReference.GetType() != "" {
			var didRegisterCompletion bool
			completionFunc := resourceNameCompletionFunc
			if field.IsList() {
				completionFunc = resourceNameListCompletionFunc
			}
			aipreflect.RangeResourceDescriptorsInPackage(
				protoregistry.GlobalFiles,
				field.ParentFile().Package(),
				func(resource *annotations.ResourceDescriptor) bool {
					if resource.GetType() == resourceReference.GetType() && len(resource.GetPattern()) > 0 {
						_ = cmd.RegisterFlagCompletionFunc(
							flag.Name,
							completionFunc(comment, resource.GetPattern()...),
						)
						didRegisterCompletion = true
						return false
					}
					return true
				},
			)
			if didRegisterCompletion {
				return
			}
		}
	}
	if field.Kind() == protoreflect.MessageKind && field.Message().FullName() == "google.protobuf.Timestamp" {
		_ = cmd.RegisterFlagCompletionFunc(flag.Name, timestampCompletionFunc(comment))
		return
	}
	if field.Kind() == protoreflect.EnumKind && !field.IsList() {
		_ = cmd.RegisterFlagCompletionFunc(flag.Name, enumFieldCompletionFunc(comment, field.Enum().Values()))
		return
	}
	// default: register active help with field comment
	_ = cmd.RegisterFlagCompletionFunc(flag.Name, fieldCompletionFunc(comment))
}

func hideImmutableForUpdateMethods(
	cmd *cobra.Command,
	flag *pflag.Flag,
	field protoreflect.FieldDescriptor,
) {
	if !isMethodType(cmd, "Update") {
		return
	}
	if fieldBehaviors, ok := proto.GetExtension(
		field.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior); ok {
		for _, fieldBehavior := range fieldBehaviors {
			if fieldBehavior == annotations.FieldBehavior_IMMUTABLE {
				_ = cmd.Flags().MarkHidden(flag.Name)
			}
		}
	}
}

func hideNameForCreateMethods(cmd *cobra.Command, flag *pflag.Flag, field protoreflect.FieldDescriptor) {
	if isMethodType(cmd, "Create") && field.Name() == FieldResourceName {
		_ = cmd.Flags().MarkHidden(flag.Name)
	}
}

func hideETagForCreateMethods(cmd *cobra.Command, flag *pflag.Flag, field protoreflect.FieldDescriptor) {
	if isMethodType(cmd, "Create") && field.Name() == "etag" {
		_ = cmd.Flags().MarkHidden(flag.Name)
	}
}

func flagName(field protoreflect.FieldDescriptor, parentFields []protoreflect.FieldDescriptor) string {
	var result strings.Builder
	for _, parentField := range parentFields {
		_, _ = result.WriteString(string(parentField.Name()))
		_ = result.WriteByte('.')
	}
	_, _ = result.WriteString(string(field.Name()))
	return strings.ReplaceAll(result.String(), "_", "-")
}

func isMethodType(cmd *cobra.Command, methodType string) bool {
	return strings.HasPrefix(string(protoreflect.FullName(cmd.Annotations[methodAnnotation]).Name()), methodType)
}

func isOutputOnly(field protoreflect.FieldDescriptor) bool {
	if fieldBehaviors, ok := proto.GetExtension(
		field.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior); ok {
		for _, fieldBehavior := range fieldBehaviors {
			if fieldBehavior == annotations.FieldBehavior_OUTPUT_ONLY {
				return true
			}
		}
	}
	return false
}

func anyParentIsOptional(parentFields []protoreflect.FieldDescriptor) bool {
	for _, field := range parentFields {
		if isOptional(field) {
			return true
		}
	}
	return false
}

func isOptional(field protoreflect.FieldDescriptor) bool {
	if fieldBehaviors, ok := proto.GetExtension(
		field.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior); ok {
		for _, fieldBehavior := range fieldBehaviors {
			if fieldBehavior == annotations.FieldBehavior_OPTIONAL {
				return true
			}
		}
	}
	return false
}
