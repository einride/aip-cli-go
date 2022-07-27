package aipcli

import (
	"context"
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Logf(ctx context.Context, format string, args ...interface{}) {
	if !ConfigFromContext(ctx).Runtime.Verbose {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, ">> %s\n", fmt.Sprintf(format, args...))
}

func logRequest(ctx context.Context, request proto.Message) {
	if !ConfigFromContext(ctx).Runtime.Verbose {
		return
	}
	for _, line := range strings.Split(protojson.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}.Format(request), "\n") {
		_, _ = fmt.Fprintf(os.Stderr, ">> %s\n", line)
	}
}

func logResponse(ctx context.Context, request proto.Message) {
	fmt.Println(protojson.MarshalOptions{Multiline: true}.Format(request))
}

func logError(ctx context.Context, err error) {
	_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
	for _, detail := range status.Convert(err).Details() {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", detail)
	}
}
