package aipcli

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stoewer/go-strcase"
	"go.einride.tech/aip/reflect/aipreflect"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// NewServiceCommand initializes a new *cobra.Command for the provided gRPC service.
func NewServiceCommand(
	service protoreflect.ServiceDescriptor,
	comments map[protoreflect.FullName]string,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   serviceUse(service),
		Short: initialUpperCase(trimComment(comments[service.FullName()])),
		Long:  comments[service.FullName()],
	}
	return cmd
}

// NewMethodCommand initializes a new *cobra.Command for the provided gRPC method.
func NewMethodCommand(
	method protoreflect.MethodDescriptor,
	in proto.Message,
	out proto.Message,
	comments map[protoreflect.FullName]string,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   methodUse(method),
		Short: initialUpperCase(trimComment(comments[method.FullName()])),
		Long:  comments[method.FullName()],
	}
	fromFile := cmd.Flags().StringP("from-file", "f", "", "path to a JSON file containing the request payload")
	_ = cmd.MarkFlagFilename("from-file", "json")
	setFlags(comments, cmd, nil, in.ProtoReflect().Descriptor(), in.ProtoReflect)
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("from-file") {
			data, err := os.ReadFile(*fromFile)
			if err != nil {
				return err
			}
			if err := protojson.Unmarshal(data, in); err != nil {
				return err
			}
		}
		conn, err := dial(cmd.Context())
		if err != nil {
			return err
		}
		LogRequest(cmd.Context(), in)
		if err := conn.Invoke(cmd.Context(), methodURI(method), in, out); err != nil {
			LogError(cmd.Context(), err)
			os.Exit(1)
		}
		LogResponse(cmd.Context(), out)
		return nil
	}
	return cmd
}

func serviceUse(service protoreflect.ServiceDescriptor) string {
	result := string(service.Name())
	result = strings.TrimSuffix(result, "Service")
	result = strcase.KebabCase(result)
	return result
}

func methodUse(method protoreflect.MethodDescriptor) string {
	result := string(method.Name())
	result = strcase.KebabCase(result)
	return result
}

func methodURI(method protoreflect.MethodDescriptor) string {
	return "/" +
		string(method.Parent().(protoreflect.ServiceDescriptor).FullName()) +
		"/" + string(method.Name())
}

func setFlags(
	comments map[protoreflect.FullName]string,
	cmd *cobra.Command,
	parentFields []protoreflect.FieldDescriptor,
	msg protoreflect.MessageDescriptor,
	mutable func() protoreflect.Message,
) {
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		switch field.Kind() {
		case protoreflect.MessageKind:
			switch field.Message().FullName() {
			case "google.protobuf.Duration":
				if field.IsList() {
					// TODO: Implement support for repeated durations.
				} else {
					addFlag(cmd, field, parentFields, comments[field.FullName()], durationValue{
						mutable: mutable,
						field:   field,
					})
				}
			case "google.protobuf.Timestamp":
				if field.IsList() {
					// TODO: Implement support for repeated timestamps.
				} else {
					addFlag(cmd, field, parentFields, comments[field.FullName()], timestampValue{
						mutable: mutable,
						field:   field,
					})
				}
			case "google.protobuf.FieldMask":
				if field.IsList() {
					// Repeated field masks is intentionally not supported.
				} else {
					addFlag(cmd, field, parentFields, comments[field.FullName()], fieldMaskValue{
						mutable: mutable,
						field:   field,
					})
				}
			default:
				if field.Cardinality() != protoreflect.Repeated {
					setFlags(
						comments,
						cmd,
						append(parentFields, field),
						field.Message(),
						func() protoreflect.Message {
							return mutable().Mutable(field).Message()
						},
					)
				}
			}
		case protoreflect.StringKind, protoreflect.BoolKind, protoreflect.BytesKind, protoreflect.DoubleKind,
			protoreflect.FloatKind, protoreflect.Int64Kind, protoreflect.Int32Kind:
			setPrimitiveFlag(comments, cmd, parentFields, mutable, field)
		}
	}
}

func setPrimitiveFlag(
	comments map[protoreflect.FullName]string,
	cmd *cobra.Command,
	parentFields []protoreflect.FieldDescriptor,
	mutable func() protoreflect.Message,
	field protoreflect.FieldDescriptor,
) {
	var value pflag.Value
	switch field.Kind() {
	case protoreflect.BoolKind:
		if field.IsList() {
			value = newPrimitiveListValue[bool](mutable, field, protoreflect.ValueOfBool, strconv.ParseBool)
		} else {
			value = newPrimitiveValue[bool](mutable, field, protoreflect.ValueOfBool, strconv.ParseBool)
		}
	case protoreflect.StringKind:
		parser := func(s string) (string, error) {
			return s, nil
		}
		if field.IsList() {
			value = newPrimitiveListValue[string](mutable, field, protoreflect.ValueOfString, parser)
		} else {
			value = newPrimitiveValue[string](mutable, field, protoreflect.ValueOfString, parser)
		}
	case protoreflect.BytesKind:
		value = newPrimitiveValue[[]byte](mutable, field, protoreflect.ValueOfBytes, base64.URLEncoding.DecodeString)
	case protoreflect.DoubleKind:
		parser := func(s string) (float64, error) {
			return strconv.ParseFloat(s, 64)
		}
		if field.IsList() {
			value = newPrimitiveListValue[float64](mutable, field, protoreflect.ValueOfFloat64, parser)
		} else {
			value = newPrimitiveValue[float64](mutable, field, protoreflect.ValueOfFloat64, parser)
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
			value = newPrimitiveListValue[float32](mutable, field, protoreflect.ValueOfFloat32, parser)
		} else {
			value = newPrimitiveValue[float32](mutable, field, protoreflect.ValueOfFloat32, parser)
		}
	case protoreflect.Int64Kind:
		parser := func(s string) (int64, error) {
			return strconv.ParseInt(s, 10, 64)
		}
		if field.IsList() {
			value = newPrimitiveListValue[int64](mutable, field, protoreflect.ValueOfInt64, parser)
		} else {
			value = newPrimitiveValue[int64](mutable, field, protoreflect.ValueOfInt64, parser)
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
			value = newPrimitiveListValue[int32](mutable, field, protoreflect.ValueOfInt32, parser)
		} else {
			value = newPrimitiveValue[int32](mutable, field, protoreflect.ValueOfInt32, parser)
		}
	default:
		panic(fmt.Errorf("unhandled primitive kind: %v", field.Kind())) // shouldn't happen
	}
	addFlag(cmd, field, parentFields, comments[field.FullName()], value)
}

func addFlag(
	cmd *cobra.Command,
	field protoreflect.FieldDescriptor,
	parentFields []protoreflect.FieldDescriptor,
	comment string,
	value pflag.Value,
) {
	flag := &pflag.Flag{
		Name:  flagName(field, parentFields),
		Usage: trimComment(comment),
		Value: value,
	}
	// Hide output only field flags.
	if fieldBehaviors, ok := proto.GetExtension(
		field.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior); ok {
		for _, fieldBehavior := range fieldBehaviors {
			if fieldBehavior == annotations.FieldBehavior_OUTPUT_ONLY {
				flag.Hidden = true
			}
		}
	}
	cmd.Flags().AddFlag(flag)
	// Mark flags as required.
	if fieldBehaviors, ok := proto.GetExtension(
		field.Options(),
		annotations.E_FieldBehavior,
	).([]annotations.FieldBehavior); ok {
		for _, fieldBehavior := range fieldBehaviors {
			if fieldBehavior == annotations.FieldBehavior_REQUIRED {
				_ = cmd.MarkFlagRequired(flag.Name)
			}
		}
	}
	// Register resource reference completion functions.
	if field.Kind() == protoreflect.StringKind {
		if resourceReference, ok := proto.GetExtension(
			field.Options(),
			annotations.E_ResourceReference,
		).(*annotations.ResourceReference); ok && resourceReference.GetType() != "" {
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
							completionFunc(resource.GetPattern()...),
						)
						return false
					}
					return true
				},
			)
		}
		// Register resource name completion functions.
		if !field.IsList() && field.Name() == "name" {
			if resourceDescriptor, ok := proto.GetExtension(
				field.Parent().Options(),
				annotations.E_Resource,
			).(*annotations.ResourceDescriptor); ok && resourceDescriptor.GetType() != "" {
				aipreflect.RangeResourceDescriptorsInPackage(
					protoregistry.GlobalFiles,
					field.ParentFile().Package(),
					func(resource *annotations.ResourceDescriptor) bool {
						if resource.GetType() == resourceDescriptor.GetType() && len(resource.GetPattern()) > 0 {
							_ = cmd.RegisterFlagCompletionFunc(
								flag.Name,
								resourceNameCompletionFunc(resource.GetPattern()...),
							)
							return false
						}
						return true
					},
				)
			}
		}
	}
}

func trimComment(comment string) string {
	result := comment
	// Clean up comment line breaks and whitespace.
	result = strings.ReplaceAll(result, "//", "")
	result = strings.ReplaceAll(result, "\n", " ")
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "  ", " ")
	result = strings.ReplaceAll(result, "  ", " ")
	// Cut out first sentence.
	if i := strings.IndexByte(result, '.'); i != -1 {
		result = result[:i]
	}
	// Trim manually documented field behavior.
	result = strings.TrimPrefix(result, "REQUIRED: ")
	result = strings.TrimPrefix(result, "Required: ")
	result = strings.ToLower(result)
	return result
}

func initialUpperCase(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if size == utf8.RuneError {
		return s
	}
	return string(unicode.ToUpper(r)) + s[size:]
}
