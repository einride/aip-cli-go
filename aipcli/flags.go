package aipcli

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func flagName(field protoreflect.FieldDescriptor, parentFields []protoreflect.FieldDescriptor) string {
	var result strings.Builder
	for _, parentField := range parentFields {
		_, _ = result.WriteString(string(parentField.Name()))
		_ = result.WriteByte('.')
	}
	_, _ = result.WriteString(string(field.Name()))
	return strings.ReplaceAll(result.String(), "_", "-")
}

func flagUsage(comment string) string {
	return trimComment(comment)
}

func newPrimitiveValue[T any](
	mutable func() protoreflect.Message,
	field protoreflect.FieldDescriptor,
	valueOf func(T) protoreflect.Value,
	parser func(string) (T, error),
) pflag.Value {
	return primitiveValue[T]{
		mutable: mutable,
		field:   field,
		valueOf: valueOf,
		parser:  parser,
	}
}

type primitiveValue[T any] struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
	valueOf func(T) protoreflect.Value
	parser  func(string) (T, error)
}

func (v primitiveValue[T]) String() string {
	return ""
}

func (v primitiveValue[T]) Set(s string) error {
	parsed, err := v.parser(s)
	if err != nil {
		return err
	}
	v.mutable().Set(v.field, v.valueOf(parsed))
	return nil
}

func (v primitiveValue[T]) Type() string {
	return v.field.Kind().String()
}

func newPrimitiveListValue[T any](
	mutable func() protoreflect.Message,
	field protoreflect.FieldDescriptor,
	valueOf func(T) protoreflect.Value,
	parser func(string) (T, error),
) pflag.Value {
	return primitiveListValue[T]{
		mutable: mutable,
		field:   field,
		parser:  parser,
		valueOf: valueOf,
	}
}

type primitiveListValue[T any] struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
	parser  func(string) (T, error)
	valueOf func(T) protoreflect.Value
}

func (v primitiveListValue[T]) String() string {
	return ""
}

func (v primitiveListValue[T]) Set(s string) error {
	values := strings.Split(s, ",")
	if len(values) == 0 {
		return nil
	}
	msg := v.mutable()
	list := msg.NewField(v.field).List()
	for _, value := range values {
		parsed, err := v.parser(value)
		if err != nil {
			return err
		}
		list.Append(v.valueOf(parsed))
	}
	msg.Set(v.field, protoreflect.ValueOfList(list))
	return nil
}

func (v primitiveListValue[T]) Type() string {
	return "[" + v.field.Kind().String() + "]"
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

type timestampValue struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
}

func (v timestampValue) String() string {
	return ""
}

func (v timestampValue) Type() string {
	return "timestamp"
}

func (v timestampValue) Set(s string) error {
	if strings.HasPrefix(s, "=") {
		return v.setExpression(strings.TrimPrefix(s, "="))
	}
	timestamp, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	v.mutable().Set(v.field, protoreflect.ValueOf(timestamppb.New(timestamp).ProtoReflect()))
	return nil
}

func (v timestampValue) setExpression(s string) error {
	env, err := cel.NewEnv(
		cel.Function(
			"now",
			cel.Overload(
				"now",
				nil,
				cel.TimestampType,
				cel.FunctionBinding(func(...ref.Val) ref.Val {
					return types.Timestamp{Time: time.Now()}
				}),
			),
		),
	)
	if err != nil {
		return err
	}
	ast, issues := env.Compile(s)
	if issues.Err() != nil {
		return issues.Err()
	}
	if ast.OutputType() != cel.TimestampType {
		return fmt.Errorf("expected timestamp but got %v", ast.OutputType())
	}
	program, err := env.Program(ast)
	if err != nil {
		return err
	}
	value, _, err := program.Eval(map[string]interface{}{})
	if err != nil {
		return err
	}
	timestamp, ok := value.Value().(time.Time)
	if !ok {
		return fmt.Errorf("invalid timestamp")
	}
	v.mutable().Set(v.field, protoreflect.ValueOf(timestamppb.New(timestamp).ProtoReflect()))
	return nil
}

type fieldMaskValue struct {
	mutable func() protoreflect.Message
	field   protoreflect.FieldDescriptor
}

func (v fieldMaskValue) String() string {
	return ""
}

func (v fieldMaskValue) Type() string {
	return "field-mask"
}

func (v fieldMaskValue) Set(s string) error {
	values := strings.Split(s, ",")
	if len(values) == 0 {
		return nil
	}
	fieldMask := fieldmaskpb.FieldMask{
		Paths: make([]string, 0, len(values)),
	}
	for _, value := range values {
		fieldMask.Paths = append(fieldMask.Paths, strings.TrimSpace(value))
	}
	v.mutable().Set(v.field, protoreflect.ValueOf(fieldMask.ProtoReflect()))
	return nil
}
