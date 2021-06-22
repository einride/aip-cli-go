package genctl

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type env struct {
	g         *protogen.GeneratedFile
	method    *protogen.Method
	preceding []*protogen.Field
	field     *protogen.Field
}

func (e *env) generate() {
	for _, annotation := range proto.GetExtension(e.field.Desc.Options(), annotations.E_FieldBehavior).([]annotations.FieldBehavior) {
		if annotation == annotations.FieldBehavior_OUTPUT_ONLY {
			return
		}
	}
	if e.field.Oneof != nil {
		e.g.P("// TODO: one-of: ", e.field.GoName)
		return
	}
	if e.field.Desc.IsList() {
		switch e.field.Desc.Kind() {
		case protoreflect.StringKind:
			e.g.P(methodCommandVariableGoName(e.method), ".Flags().StringSliceVar(&", e.structPath(), ", ", strconv.Quote(e.fieldPath()), ", []string{}, ", strconv.Quote(e.comment()), ")")
		default:
			e.g.P("// TODO: list: ", e.field.GoName, " ", e.field.Desc.Kind())
		}
		return
	}
	if e.field.Desc.IsMap() {
		e.g.P("// TODO: map: ", e.field.GoName, e.field.Desc.MapKey().Kind(), " ", e.field.Desc.Kind())
		return
	}
	switch e.field.Desc.Kind() {
	case protoreflect.MessageKind:
		e.g.P(e.structPath(), " = new(", e.field.Message.GoIdent, ")")
		for _, nested := range e.field.Message.Fields {
			ne := env{
				g:         e.g,
				method:    e.method,
				preceding: append(e.preceding, e.field),
				field:     nested,
			}
			ne.generate()
		}
	case protoreflect.StringKind:
		e.g.P(methodCommandVariableGoName(e.method), ".Flags().StringVar(&", e.structPath(), ", ", strconv.Quote(e.fieldPath()), ", \"\", ", strconv.Quote(e.comment()), ")")
	case protoreflect.BoolKind:
		e.g.P(methodCommandVariableGoName(e.method), ".Flags().BoolVar(&", e.structPath(), ", ", strconv.Quote(e.fieldPath()), ", false, ", strconv.Quote(e.comment()), ")")
	case protoreflect.Int32Kind:
		e.g.P(methodCommandVariableGoName(e.method), ".Flags().Int32Var(&", e.structPath(), ", ", strconv.Quote(e.fieldPath()), ", 10, ", strconv.Quote(e.comment()), ")")
	case protoreflect.Int64Kind:
		e.g.P(methodCommandVariableGoName(e.method), ".Flags().Int64Var(&", e.structPath(), ", ", strconv.Quote(e.fieldPath()), ", 10, ", strconv.Quote(e.comment()), ")")
	case protoreflect.FloatKind:
		e.g.P(methodCommandVariableGoName(e.method), ".Flags().Float32Var(&", e.structPath(), ", ", strconv.Quote(e.fieldPath()), ", 10, ", strconv.Quote(e.comment()), ")")
	case protoreflect.DoubleKind:
		e.g.P(methodCommandVariableGoName(e.method), ".Flags().Float64Var(&", e.structPath(), ", ", strconv.Quote(e.fieldPath()), ", 10, ", strconv.Quote(e.comment()), ")")
	default:
		e.g.P("// TODO: ", e.field.Desc.Kind(), " ", e.field.GoName)
	}
}

func (e *env) structPath() string {
	fields := make([]string, 0, len(e.preceding)+1)
	for _, f := range append(e.preceding, e.field) {
		fields = append(fields, f.GoName)
	}
	return fmt.Sprintf("%s.%s", requestVariableGoName(e.method), strings.Join(fields, "."))
}

func (e *env) fieldPath() string {
	fields := make([]string, 0, len(e.preceding)+1)
	for _, f := range append(e.preceding, e.field) {
		fields = append(fields, f.Desc.JSONName())
	}
	return strings.Join(fields, ".")
}

func (e *env) comment() string {
	s := strings.ReplaceAll(string(e.field.Comments.Leading), "\n ", "\n")
	return strings.TrimSpace(s)
}
