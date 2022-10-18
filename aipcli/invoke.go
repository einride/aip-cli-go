package aipcli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/itchyny/gojq"
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
		for _, line := range strings.Split(marshalJson(request), "\n") {
			cmd.PrintErrln(">>", line)
		}
	}
	if err := conn.Invoke(cmd.Context(), uri, request, response); err != nil {
		for _, detail := range status.Convert(err).Details() {
			cmd.PrintErrln(detail)
		}
		return err // let Cobra print the error when exiting
	}
	out, err := formatOutput(marshalJson(response), cmd)
	if err != nil {
		return fmt.Errorf("unable to format output: %v", err)
	}
	cmd.Print(out)
	return nil
}

func marshalJson(msg proto.Message) string {
	return protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		EmitUnpopulated: true,
	}.Format(msg)
}

func formatOutput(in string, cmd *cobra.Command) (string, error) {
	format := getFormat(cmd)
	query, err := gojq.Parse(format)
	if err != nil {
		return "", err
	}
	var input map[string]interface{}
	if err := json.Unmarshal([]byte(in), &input); err != nil {
		return "", err
	}
	iter := query.Run(input)
	var builder strings.Builder
	for i := 0; ; i++ {
		v, ok := iter.Next()
		if !ok {
			break
		} else if builder.Len() > 0 {
			builder.WriteByte('\n')
		}
		if err, ok := v.(error); ok {
			return "", err
		}

		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return "", err
		}
		builder.WriteString(string(b))
	}
	return builder.String(), nil
}
