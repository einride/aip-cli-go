package plugin

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/einride/ctl/internal/codegen"
	"github.com/stoewer/go-strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldGenerator struct {
	method    protoreflect.MethodDescriptor
	preceding []protoreflect.FieldDescriptor
	message   protoreflect.MessageDescriptor
}

func (fi FieldGenerator) Generate(f *codegen.File) {
	for k := 0; k < fi.message.Fields().Len(); k++ {
		field := fi.message.Fields().Get(k)
		fi.generateField(f, field)
	}
}

func (fi FieldGenerator) jsonPath(field protoreflect.FieldDescriptor) string {
	p := make([]string, 0, len(fi.preceding)+1)
	for _, f := range append(fi.preceding, field) {
		p = append(p, f.JSONName())
	}
	return strings.Join(p, ".")
}

func (fi FieldGenerator) structPath(field protoreflect.FieldDescriptor) string {
	p := make([]string, 0, len(fi.preceding)+1)
	for _, f := range append(fi.preceding, field) {
		p = append(p, strcase.UpperCamelCase(string(f.Name())))
	}
	return strings.Join(p, ".")
}

func (fi FieldGenerator) msgName(msg protoreflect.MessageDescriptor) string {
	name := string(msg.Name())
	for {
		parentMsg, ok := msg.Parent().(protoreflect.MessageDescriptor)
		if !ok {
			break
		}
		name = string(parentMsg.Name()) + "_" + name
		msg = parentMsg
	}
	return name
}

func (fi FieldGenerator) generateField(f *codegen.File, field protoreflect.FieldDescriptor) {
	// todo support nested
	fieldName := fi.jsonPath(field)
	structField := fmt.Sprintf("%s.%s", methodInputVarName(fi.method), fi.structPath(field))
	comment := field.ParentFile().SourceLocations().ByDescriptor(field).LeadingComments

	if field.ContainingOneof() != nil {
		f.P("// TODO: one-of: ", fieldName)
		return
	}

	if field.IsList() {
		switch field.Kind() {
		case protoreflect.StringKind:
			f.P(methodCmdVarName(fi.method), ".Flags().StringSliceVar(&", structField, ", ", strconv.Quote(fieldName), ", []string{}, ", strconv.Quote(comment), ")")
		default:
			f.P("// TODO: list: ", fieldName, " ", field.Kind())
		}
		return
	}

	if field.IsMap() {
		f.P("// TODO: map: ", fieldName, field.MapKey().Kind(), " ", field.Kind())
		return
	}

	switch field.Kind() {
	case protoreflect.MessageKind:
		msg := field.Message()
		pbPkg := f.Import(getGoPkg(msg).path)
		f.P(structField, " = new(", pbPkg, ".", fi.msgName(msg), ")")
		FieldGenerator{
			method:    fi.method,
			preceding: append(fi.preceding, field),
			message:   msg,
		}.Generate(f)
	case protoreflect.StringKind:
		f.P(methodCmdVarName(fi.method), ".Flags().StringVar(&", structField, ", ", strconv.Quote(fieldName), ", \"\", ", strconv.Quote(comment), ")")
	case protoreflect.BoolKind:
		f.P(methodCmdVarName(fi.method), ".Flags().BoolVar(&", structField, ", ", strconv.Quote(fieldName), ", false, ", strconv.Quote(comment), ")")
	case protoreflect.Int32Kind:
		f.P(methodCmdVarName(fi.method), ".Flags().Int32Var(&", structField, ", ", strconv.Quote(fieldName), ", 0, ", strconv.Quote(comment), ")")
	case protoreflect.Int64Kind:
		f.P(methodCmdVarName(fi.method), ".Flags().Int64Var(&", structField, ", ", strconv.Quote(fieldName), ", 0, ", strconv.Quote(comment), ")")
	case protoreflect.FloatKind:
		f.P(methodCmdVarName(fi.method), ".Flags().Float32Var(&", structField, ", ", strconv.Quote(fieldName), ", 0, ", strconv.Quote(comment), ")")
	case protoreflect.DoubleKind:
		f.P(methodCmdVarName(fi.method), ".Flags().Float64Var(&", structField, ", ", strconv.Quote(fieldName), ", 0, ", strconv.Quote(comment), ")")
	default:
		f.P("// TODO: ", field.Kind(), " ", field.Name())
	}
}
