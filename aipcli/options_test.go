package aipcli

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"gotest.tools/v3/assert"
)

func TestWithGlobalTokenFunc_setsTokenFuncInOptions(t *testing.T) {
	var called bool
	var calledCmd *cobra.Command
	opt := WithGlobalTokenFunc(func(cmd *cobra.Command) (string, error) {
		called = true
		calledCmd = cmd
		return "abc", nil
	})
	opts := &Options{}
	opt(opts)
	assert.Assert(t, opts.tokenFunc != nil)
	token, err := opts.tokenFunc(&cobra.Command{Use: "foobar"})
	assert.NilError(t, err)
	assert.Equal(t, token, "abc")
	assert.Assert(t, called)
	assert.Assert(t, calledCmd != nil)
}

func TestDefaultTokenFunc_behavior(t *testing.T) {
	t.Run("returns error with missing file if CachedIdentityTokenPath set", func(t *testing.T) {
		cmd := &cobra.Command{}
		initContext(cmd, Config{CachedIdentityTokenPath: "not_a_file.json"})
		token, err := defaultTokenFunc(cmd)
		assert.ErrorContains(t, err, "failed to read identity token from config file")
		assert.Equal(t, token, "")
	})

	t.Run("returns real token or non-empty string for GoogleCloudIdentityTokens", func(t *testing.T) {
		cmd := &cobra.Command{}
		initContext(cmd, Config{GoogleCloudIdentityTokens: true})
		token, err := defaultTokenFunc(cmd)
		assert.NilError(t, err)
		assert.Assert(t, token != "")
	})

	t.Run("returns empty string if config blank", func(t *testing.T) {
		cmd := &cobra.Command{}
		initContext(cmd, Config{})
		token, err := defaultTokenFunc(cmd)
		assert.NilError(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("returns empty string if config not initialized", func(t *testing.T) {
		cmd := &cobra.Command{}
		token, err := defaultTokenFunc(cmd)
		assert.NilError(t, err)
		assert.Equal(t, token, "")
	})
}

func Test_getOptions(t *testing.T) {
	t.Run("root context with options returns them", func(t *testing.T) {
		cmd := &cobra.Command{}
		want := "tokroot"
		setCommandOptions(cmd, Options{
			tokenFunc: func(*cobra.Command) (string, error) { return want, nil },
		})
		opts := getOptions(cmd)
		assert.Assert(t, opts.tokenFunc != nil)
		token, err := opts.tokenFunc(cmd)
		assert.NilError(t, err)
		assert.Equal(t, token, want)
	})
	t.Run("child gets root options", func(t *testing.T) {
		root := &cobra.Command{}
		child := &cobra.Command{}
		root.AddCommand(child)
		want := "tokfromroot"
		setCommandOptions(root, Options{
			tokenFunc: func(*cobra.Command) (string, error) { return want, nil },
		})
		opts := getOptions(child)
		assert.Assert(t, opts.tokenFunc != nil)
		token, err := opts.tokenFunc(child)
		assert.NilError(t, err)
		assert.Equal(t, token, want)
	})
	t.Run("grandchild gets root options", func(t *testing.T) {
		root := &cobra.Command{}
		parent := &cobra.Command{}
		child := &cobra.Command{}
		root.AddCommand(parent)
		parent.AddCommand(child)
		want := "tokgrandroot"
		setCommandOptions(root, Options{
			tokenFunc: func(*cobra.Command) (string, error) { return want, nil },
		})
		opts := getOptions(child)
		assert.Assert(t, opts.tokenFunc != nil)
		token, err := opts.tokenFunc(child)
		assert.NilError(t, err)
		assert.Equal(t, token, want)
	})
	t.Run("no options if root has no context at all", func(t *testing.T) {
		cmd := &cobra.Command{}
		opts := getOptions(cmd)
		assert.Assert(t, opts.tokenFunc == nil)
	})
	t.Run("no options if root context has wrong key", func(t *testing.T) {
		type otherKey struct{}
		cmd := &cobra.Command{}
		cmd.SetContext(context.WithValue(context.Background(), otherKey{}, "x"))
		opts := getOptions(cmd)
		assert.Assert(t, opts.tokenFunc == nil)
	})
	t.Run("handles nil context on root", func(t *testing.T) {
		cmd := &cobra.Command{}
		opts := getOptions(cmd)
		assert.Assert(t, opts.tokenFunc == nil)
	})
	t.Run("handles child when root nil context", func(t *testing.T) {
		root := &cobra.Command{}
		child := &cobra.Command{}
		root.AddCommand(child)
		opts := getOptions(child)
		assert.Assert(t, opts.tokenFunc == nil)
	})
}

func Test_collectChildrenOptions(t *testing.T) {
	t.Run("finds options from the root", func(t *testing.T) {
		cmd := &cobra.Command{}
		expected := "tok"
		setCommandOptions(cmd, Options{
			tokenFunc: func(*cobra.Command) (string, error) { return expected, nil },
		})
		opts := collectChildrenOptions(cmd)
		assert.Assert(t, opts.tokenFunc != nil)
		value, err := opts.tokenFunc(cmd)
		assert.NilError(t, err)
		assert.Equal(t, value, expected)
	})
	t.Run("finds options from child", func(t *testing.T) {
		root := &cobra.Command{}
		root.SetContext(context.Background())
		child := &cobra.Command{}
		root.AddCommand(child)
		expected := "tokchild"
		setCommandOptions(child, Options{
			tokenFunc: func(*cobra.Command) (string, error) { return expected, nil },
		})
		opts := collectChildrenOptions(root)
		assert.Assert(t, opts.tokenFunc != nil)
		value, err := opts.tokenFunc(child)
		assert.NilError(t, err)
		assert.Equal(t, value, expected)
	})
	t.Run("finds options from grandchild", func(t *testing.T) {
		root := &cobra.Command{}
		parent := &cobra.Command{}
		child := &cobra.Command{}
		root.SetContext(context.Background())
		parent.SetContext(context.Background())
		root.AddCommand(parent)
		parent.AddCommand(child)
		expected := "tokgc"
		setCommandOptions(child, Options{
			tokenFunc: func(*cobra.Command) (string, error) { return expected, nil },
		})
		opts := collectChildrenOptions(root)
		assert.Assert(t, opts.tokenFunc != nil)
		val, err := opts.tokenFunc(child)
		assert.NilError(t, err)
		assert.Equal(t, val, expected)
	})
	t.Run("warns and uses first found when multiple children define tokenFunc", func(t *testing.T) {
		root := &cobra.Command{Use: "root"}
		child1 := &cobra.Command{Use: "child1"}
		child2 := &cobra.Command{Use: "child2"}
		root.AddCommand(child1, child2)
		root.SetContext(context.Background())
		setCommandOptions(child1, Options{
			tokenFunc: func(*cobra.Command) (string, error) { return "one", nil },
		})
		setCommandOptions(child2, Options{
			tokenFunc: func(*cobra.Command) (string, error) { return "two", nil },
		})

		r, w, err := os.Pipe()
		assert.NilError(t, err)
		child2.SetErr(w)

		opts := collectChildrenOptions(root)
		w.Close()

		assert.Assert(t, opts.tokenFunc != nil)
		token, err := opts.tokenFunc(child1)
		assert.NilError(t, err)
		assert.Equal(t, token, "one")

		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		assert.NilError(t, err)
		out := buf.String()
		assert.Assert(t, strings.Contains(out, "WARNING"), out)
		assert.Assert(t, strings.Contains(out, "child2 attempted to configure a global token function"), out)
	})
	t.Run("returns empty options if no children or root options found", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.SetContext(context.Background())
		opts := collectChildrenOptions(cmd)
		assert.Assert(t, opts.tokenFunc == nil)
	})
}

func Test_getContextOptions(t *testing.T) {
	t.Run("extracts options from context with valid value", func(t *testing.T) {
		expect := "mmm"
		ctx := context.WithValue(context.Background(), optionsContextKey{}, &optionsContextValue{
			options: Options{tokenFunc: func(*cobra.Command) (string, error) { return expect, nil }},
		})
		opts := getContextOptions(ctx)
		assert.Assert(t, opts.tokenFunc != nil)
		tok, err := opts.tokenFunc(&cobra.Command{})
		assert.NilError(t, err)
		assert.Equal(t, tok, expect)
	})
	t.Run("returns empty options for context with missing key", func(t *testing.T) {
		opts := getContextOptions(context.Background())
		assert.Assert(t, opts.tokenFunc == nil)
	})
	t.Run("returns empty options for context with wrong type", func(t *testing.T) {
		type wrongKey struct{}
		ctx := context.WithValue(context.Background(), wrongKey{}, "not-options-value")
		opts := getContextOptions(ctx)
		assert.Assert(t, opts.tokenFunc == nil)
	})
}
