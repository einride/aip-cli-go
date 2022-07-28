package aipcli

import (
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func invoke(cmd *cobra.Command, uri string, request, response proto.Message) error {
	conn, err := dial(cmd)
	if err != nil {
		return err
	}
	if isVerbose(cmd) {
		for _, line := range strings.Split(format(request), "\n") {
			cmd.PrintErrln(">>", line)
		}
	}
	if err := conn.Invoke(cmd.Context(), uri, request, response); err != nil {
		for _, detail := range status.Convert(err).Details() {
			cmd.PrintErrln(detail)
		}
		return err // let Cobra print the error when exiting
	}
	cmd.Println(format(response))
	return nil
}

func format(msg proto.Message) string {
	return protojson.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}.Format(msg)
}
