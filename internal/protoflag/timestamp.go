package protovalue

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Timestamp(mutable func() protoreflect.Message, field protoreflect.FieldDescriptor) pflag.Value {
	return timestampValue{mutable: mutable, field: field}
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
