package aipcli

import (
	"testing"

	"github.com/spf13/cobra"
	testservicefreightv1 "go.einride.tech/aip-cli/aipcli/testservice"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
)

func Test_resetOptionalFlags(t *testing.T) {
	t.Parallel()
	cmd := &cobra.Command{}
	var (
		pageSize        int32
		pageToken       string
		myOptionalField string
	)
	cmd.Flags().Int32Var(&pageSize, "page-size", 0, "")
	cmd.Flags().StringVar(&pageToken, "page-token", "", "")
	cmd.Flags().StringVar(&myOptionalField, "my-optional-field", "", "")

	// Simulate "my-optional-field" being annotated as a proto value.
	err := cmd.Flags().SetAnnotation(
		"my-optional-field",
		fieldNameAnnotation,
		[]string{"einride.example.freight.v1.ListShippersRequest.my_optional_field"},
	)
	assert.NilError(t, err)

	t.Run("optional not provided", func(t *testing.T) {
		t.Parallel()
		request := &testservicefreightv1.ListShippersRequest{}
		resetOptionalFlags(cmd, request)
		assert.Assert(t, request.MyOptionalField == nil)
	})

	for _, tt := range []struct {
		name               string
		optionalFieldValue string
	}{
		{
			name:               "optional provided - empty string",
			optionalFieldValue: "",
		},
		{
			name:               "optional provided - non-empty string",
			optionalFieldValue: "value",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			request := &testservicefreightv1.ListShippersRequest{
				MyOptionalField: proto.String(tt.optionalFieldValue),
			}
			assert.NilError(t, cmd.Flags().Set("my-optional-field", tt.optionalFieldValue))
			resetOptionalFlags(cmd, request)
			assert.Equal(t, tt.optionalFieldValue, request.GetMyOptionalField())
		})
	}
}

func Test_qualifiedServiceUse(t *testing.T) {
	for _, tt := range []struct {
		name     string
		expected []string
	}{
		{
			name: "google.iam.v1.IAMPolicy",
			expected: []string{
				"iam-policy",
				"iam-policy-v1",
				"iam-policy-v1-iam",
				"iam-policy-v1-iam-google",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual := make([]string, 0, len(tt.expected))
			var i int
			for {
				result, ok := qualifiedServiceUse(protoreflect.FullName(tt.name), i)
				if !ok {
					break
				}
				actual = append(actual, result)
				i++
			}
			assert.DeepEqual(t, tt.expected, actual)
		})
	}
}
