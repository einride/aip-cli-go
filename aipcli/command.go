package aipcli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const fromFileFlag = "from-file"

// NewCommand initializes a new *cobra.Command with AIP CLI flags and help functions.
func NewCommand(cmd *cobra.Command, commands ...*cobra.Command) *cobra.Command {
	initPersistentFlags(cmd)
	cmd.SetHelpFunc(helpFunc)
	cmd.SetUsageFunc(usageFunc)
	deduplicateServiceCommandUses(commands)
	cmd.AddCommand(commands...)
	return cmd
}

// NewModuleCommand initializes a new *cobra.Command for a CLI module.
// A module is a collection of services with a common CLI config.
func NewModuleCommand(
	use string,
	short string,
	config Config,
	commands ...*cobra.Command,
) *cobra.Command {
	cmd := NewCommand(&cobra.Command{
		Use:   use,
		Short: short,
		Annotations: map[string]string{
			moduleAnnotation: use,
		},
	}, commands...)
	setConfig(cmd, config)
	return cmd
}

// NewServiceCommand initializes a new *cobra.Command for the provided gRPC service.
func NewServiceCommand(
	config Config,
	service protoreflect.ServiceDescriptor,
	comments map[protoreflect.FullName]string,
	commands ...*cobra.Command,
) *cobra.Command {
	cmd := NewCommand(&cobra.Command{
		Use:   serviceUse(service.FullName()),
		Short: initialUpperCase(trimFieldComment(comments[service.FullName()])),
		Long:  trimLongComment(comments[service.FullName()]),
		Annotations: map[string]string{
			serviceAnnotation: string(service.FullName()),
		},
	}, commands...)
	setConfig(cmd, config)
	return cmd
}

// NewMethodCommand initializes a new *cobra.Command for the provided gRPC method.
func NewMethodCommand(
	config Config,
	method protoreflect.MethodDescriptor,
	in proto.Message,
	out proto.Message,
	comments map[protoreflect.FullName]string,
) *cobra.Command {
	cmd := NewCommand(&cobra.Command{
		Use:   methodUse(method),
		Short: initialUpperCase(trimFieldComment(comments[method.FullName()])),
		Long:  trimLongComment(comments[method.FullName()]),
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			initContext(cmd, config)
			setDefaultHostFromProto(cmd, method)
		},
		Annotations: map[string]string{
			methodAnnotation: string(method.FullName()),
		},
	})
	setConfig(cmd, config)
	fromFile := cmd.Flags().StringP(fromFileFlag, "f", "", "path to a JSON file containing the request payload")
	_ = cmd.MarkFlagFilename(fromFileFlag, "json")
	setFlags(comments, cmd, nil, in.ProtoReflect().Descriptor(), in.ProtoReflect)
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		if cmd.Flags().Changed(fromFileFlag) {
			data, err := os.ReadFile(*fromFile)
			if err != nil {
				return err
			}
			if err := protojson.Unmarshal(data, in); err != nil {
				return err
			}
		}
		resetOptionalFlags(cmd, in)
		return invoke(cmd, methodURI(method), in, out)
	}
	return cmd
}

// Resets to nil the fields that:
// - correspond to an optional field in the protos AND
// - didn't have any value specified in the command invocation.
func resetOptionalFlags(cmd *cobra.Command, request proto.Message) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		// If a value was provided by the user, continue.
		if flag.Changed {
			return
		}
		// If the field doesn't correspond to a proto field (e.g. connection flags), continue.
		fieldNames, ok := flag.Annotations[fieldNameAnnotation]
		if !ok || len(fieldNames) == 0 {
			return
		}
		fieldName := protoreflect.FullName(fieldNames[0])
		field, err := protoregistry.GlobalFiles.FindDescriptorByName(fieldName)
		if err != nil {
			return
		}
		fieldDescriptor, ok := field.(protoreflect.FieldDescriptor)
		if !ok {
			return
		}
		// If the field is optional, clear the value.
		if fieldDescriptor.Cardinality() == protoreflect.Optional {
			request.ProtoReflect().Clear(fieldDescriptor)
		}
	})
}
