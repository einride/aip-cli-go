package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	deployctl "go.einride.tech/aip-cli/cmd/deployctl/gen"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"
)

func TestCLIHelpOutput(t *testing.T) {
	for _, args := range getCommandArgs() {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			// Arrange
			var buf bytes.Buffer
			root := newCommand()
			root.SetOut(&buf)
			root.SetErr(&buf)
			root.SetArgs(append(args, "--help"))

			// Act
			err := root.Execute()

			// Assert
			assert.NilError(t, err)
			goldenFile := fmt.Sprintf("deployctl-%s.golden", strings.Join(args, "-"))
			golden.Assert(t, buf.String(), goldenFile)
		})
	}
}

func newCommand() *cobra.Command {
	return deployctl.NewModule().Command("deployctl", "deploy CLI tool")
}

func getCommandArgs() [][]string {
	rootCommand := newCommand()
	return getArgs(nil, rootCommand)
}

// args recursively collects all sub command args.
// Example: [ [map, bach-get], [map, create], ... ].
func getArgs(cmds []string, cmd *cobra.Command) [][]string {
	children := cmd.Commands()
	if len(children) == 0 {
		return [][]string{cmds}
	}
	var result [][]string
	for _, child := range children {
		result = append(result, getArgs(append(cmds, child.Name()), child)...)
	}
	return result
}
