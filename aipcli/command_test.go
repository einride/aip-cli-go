package aipcli

import (
	"testing"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gotest.tools/v3/assert"
)

func Test_getAddress_rootFlagNotShadowedByLocalField(t *testing.T) {
	// Regression test: a proto request with a field named "address" causes the generated
	// subcommand to register a local --address flag that shadows the root persistent
	// connection flag in the merged flagset. With TraverseChildren the connection address
	// lands in the root's own persistent flags; getAddress must walk the parent chain to
	// find it rather than reading from the (shadowed) merged flagset.
	root := &cobra.Command{Use: "root", TraverseChildren: true}
	initPersistentFlags(root)
	setConfig(root, Config{})

	sub := &cobra.Command{Use: "sub"}
	root.AddCommand(sub)
	initPersistentFlags(sub) // every command level registers its own persistent flags
	// Simulate the generated local --address flag for a proto "address" string field.
	sub.Flags().String(addressFlag, "", "postal address proto field")

	// Simulate TraverseChildren parsing --address localhost:50051 at root level.
	if err := root.PersistentFlags().Set(addressFlag, "localhost:50051"); err != nil {
		t.Fatal(err)
	}

	addr, ok := getAddress(sub)
	assert.Assert(t, ok)
	assert.Equal(t, addr, "localhost:50051")
}

func Test_isInsecure_rootFlagNotShadowedByLeafPersistent(t *testing.T) {
	// Regression test: each command level registers its own --insecure persistent flag
	// via initPersistentFlags. With TraverseChildren the root-level --insecure lands on
	// root's own persistent flag, not the leaf's. isInsecure must walk the parent chain.
	root := &cobra.Command{Use: "root", TraverseChildren: true}
	initPersistentFlags(root)
	setConfig(root, Config{})

	sub := &cobra.Command{Use: "sub"}
	root.AddCommand(sub)
	initPersistentFlags(sub) // leaf has its own --insecure, not set

	// Simulate TraverseChildren parsing --insecure at root level.
	if err := root.PersistentFlags().Set(insecureFlag, "true"); err != nil {
		t.Fatal(err)
	}

	assert.Assert(t, isInsecure(sub))
}

// Test_NewMethodCommand_flagRedefinedWhenRequiredFieldPrecedesAddress is a regression
// test for the construction-time panic (PR #284). Before the fix, markRequiredFlags
// called cmd.MarkFlagsOneRequired during proto-field iteration; that triggered
// mergePersistentFlags(), which copied --address from PersistentFlags() into local
// Flags() before the colliding proto-field flag was added. The fix defers
// MarkFlagsOneRequired until after every proto-field flag is registered; pflag's
// AddFlagSet then skips the persistent --address copy because the local one already
// exists.
func Test_NewMethodCommand_flagRedefinedWhenRequiredFieldPrecedesAddress(t *testing.T) {
	cmd := &cobra.Command{Use: "get-book"}
	initPersistentFlags(cmd)
	_ = cmd.Flags().StringP(fromFileFlag, "f", "", "")
	// Register all proto-field flags first (including the colliding "address").
	_ = cmd.Flags().String("name", "", "")
	_ = cmd.Flags().String(addressFlag, "", "") // proto field "address" — must precede MarkFlagsOneRequired
	// Deferred MarkFlagsOneRequired: by now local Flags() already has --address,
	// so the subsequent mergePersistentFlags() skips the persistent copy silently.
	cmd.MarkFlagsOneRequired("name", fromFileFlag) // must not panic
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
