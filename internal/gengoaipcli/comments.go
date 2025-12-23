package gengoaipcli

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

func collectDescriptorComments(comments map[protoreflect.FullName]string, desc protoreflect.Descriptor) {
	comments[desc.FullName()] = desc.ParentFile().SourceLocations().ByDescriptor(desc).LeadingComments
}

func collectMethodComments(comments map[protoreflect.FullName]string, method protoreflect.MethodDescriptor) {
	collectDescriptorComments(comments, method)
	collectMessageComments(comments, method.Input())
}

func collectMessageComments(comments map[protoreflect.FullName]string, msg protoreflect.MessageDescriptor) {
	if _, ok := comments[msg.FullName()]; ok {
		// already collected comments for this message
		return
	}
	collectDescriptorComments(comments, msg)
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		collectDescriptorComments(comments, field)
		if field.Kind() == protoreflect.MessageKind {
			collectMessageComments(comments, field.Message())
		}
	}
}
