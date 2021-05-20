package plugin

import (
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"

	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func serviceCmdVarName(service protoreflect.ServiceDescriptor) string {
	segments := strings.Split(string(service.FullName()), ".")
	return strings.Join(segments, "_")
}

func serviceCmdName(service protoreflect.ServiceDescriptor) string {
	segments := strings.Split(string(service.FullName()), ".")
	segments[len(segments)-1] = strings.Replace(segments[len(segments)-1], "Service", "", 1)
	return strings.Join(segments, ".")
}

func methodCmdVarName(method protoreflect.MethodDescriptor) string {
	segments := strings.Split(string(method.FullName()), ".")
	return strings.Join(segments, "_")
}

func methodInputVarName(method protoreflect.MethodDescriptor) string {
	segments := strings.Split(string(method.Input().FullName()), ".")
	return strings.Join(segments, "_")
}

func methodCmdName(method protoreflect.MethodDescriptor) string {
	return string(method.Name())
}

type goPkg struct {
	path string
	name string
}

func getFileGoPkg(fileDesc *descriptorpb.FileDescriptorProto) goPkg {
	pkg := fileDesc.Options.GetGoPackage()
	if strings.Contains(pkg, ";") {
		parts := strings.Split(pkg, ";")
		return goPkg{
			path: parts[0],
			name: parts[1],
		}
	}
	return goPkg{
		path: pkg,
		name: "pb",
	}
}

func getInputPkg(meth protoreflect.MethodDescriptor) goPkg {
	return getFileGoPkg(protodesc.ToFileDescriptorProto(meth.Input().ParentFile()))
}

func getGoPkg(desc protoreflect.Descriptor) goPkg {
	return getFileGoPkg(protodesc.ToFileDescriptorProto(desc.ParentFile()))
}
