package aipcli

import (
	"context"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"gotest.tools/v3/assert"
)

func TestDefaultTokenFunc_behavior(t *testing.T) {
	t.Run("returns error with missing file if CachedIdentityTokenPath set", func(t *testing.T) {
		// Given
		cmd := &cobra.Command{}
		initContext(cmd, Config{CachedIdentityTokenPath: "not_a_file.json"})
		// When
		token, err := defaultTokenFunc(cmd)
		// Then
		assert.ErrorContains(t, err, "failed to read identity token from config file")
		assert.Equal(t, token, "")
	})

	t.Run("returns empty string if config blank", func(t *testing.T) {
		// Given
		cmd := &cobra.Command{}
		initContext(cmd, Config{})
		// When
		token, err := defaultTokenFunc(cmd)
		// Then
		assert.NilError(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("returns empty string if config not initialized", func(t *testing.T) {
		// Given
		cmd := &cobra.Command{}
		// When
		token, err := defaultTokenFunc(cmd)
		// Then
		assert.NilError(t, err)
		assert.Equal(t, token, "")
	})
}

func TestSetCommandOptions(t *testing.T) {
	t.Run("sets options when context exists", func(t *testing.T) {
		// Given
		cmd := &cobra.Command{}
		cmd.SetContext(context.WithValue(context.Background(), contextKey{}, &contextValue{}))
		opts := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "test-token", nil
			},
		}
		// When
		setCommandOptions(cmd, opts)

		// Then
		ctx := cmd.Context()
		assert.Assert(t, ctx != nil)
		ctxVal := ctx.Value(optionsContextKey{})
		assert.Assert(t, ctxVal != nil)
		typedVal, ok := ctxVal.(*optionsContextValue)
		assert.Assert(t, ok)
		assert.Assert(t, typedVal.options.tokenFunc != nil)
		token, err := typedVal.options.tokenFunc(cmd)
		assert.NilError(t, err)
		assert.Equal(t, token, "test-token")
	})

	t.Run("sets options when context is nil", func(t *testing.T) {
		// Given
		cmd := &cobra.Command{}
		opts := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "test-token-2", nil
			},
		}
		// When
		setCommandOptions(cmd, opts)

		// Then
		ctx := cmd.Context()
		assert.Assert(t, ctx != nil)
		ctxVal := ctx.Value(optionsContextKey{})
		assert.Assert(t, ctxVal != nil)
		typedVal, ok := ctxVal.(*optionsContextValue)
		assert.Assert(t, ok)
		assert.Assert(t, typedVal.options.tokenFunc != nil)
		token, err := typedVal.options.tokenFunc(cmd)
		assert.NilError(t, err)
		assert.Equal(t, token, "test-token-2")
	})

	t.Run("sets options with nil tokenFunc", func(t *testing.T) {
		// Given
		cmd := &cobra.Command{}
		opts := Options{
			tokenFunc: nil,
		}
		// When
		setCommandOptions(cmd, opts)

		// Then
		ctx := cmd.Context()
		assert.Assert(t, ctx != nil)
		ctxVal := ctx.Value(optionsContextKey{})
		assert.Assert(t, ctxVal != nil)
		typedVal, ok := ctxVal.(*optionsContextValue)
		assert.Assert(t, ok)
		assert.Assert(t, typedVal.options.tokenFunc == nil)
	})
}

func TestGetOptionFromChildren(t *testing.T) {
	t.Run("returns nil when no options found", func(t *testing.T) {
		// Given
		cmd := &cobra.Command{}
		// When
		result := getOptionFromChildren(cmd, func(currentVal any, _ Options) any {
			return currentVal
		})
		// Then
		assert.Assert(t, result == nil)
	})

	t.Run("finds options in root command", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		opts := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "root-token", nil
			},
		}
		setCommandOptions(root, opts)

		var foundTokenFunc *globalTokenFunc
		// When
		result := getOptionFromChildren(root, func(_ any, allOpts Options) any {
			if allOpts.tokenFunc != nil {
				foundTokenFunc = &allOpts.tokenFunc
			}
			return foundTokenFunc
		})
		// Then
		assert.Assert(t, result != nil)
		tokenFunc, ok := result.(*globalTokenFunc)
		assert.Assert(t, ok)
		token, err := (*tokenFunc)(root)
		assert.NilError(t, err)
		assert.Equal(t, token, "root-token")
	})

	t.Run("finds options in child command", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		child := &cobra.Command{Use: "child"}
		root.AddCommand(child)

		opts := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "child-token", nil
			},
		}
		setCommandOptions(child, opts)

		// When
		var foundTokenFunc *globalTokenFunc
		result := getOptionFromChildren(root, func(_ any, allOpts Options) any {
			if allOpts.tokenFunc != nil {
				foundTokenFunc = &allOpts.tokenFunc
			}
			return foundTokenFunc
		})
		// Then
		assert.Assert(t, result != nil)
		tokenFunc, ok := result.(*globalTokenFunc)
		assert.Assert(t, ok)
		token, err := (*tokenFunc)(child)
		assert.NilError(t, err)
		assert.Equal(t, token, "child-token")
	})

	t.Run("finds options in nested child commands", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		child1 := &cobra.Command{Use: "child1"}
		child2 := &cobra.Command{Use: "child2"}
		root.AddCommand(child1)
		child1.AddCommand(child2)

		opts := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "child2-token", nil
			},
		}
		setCommandOptions(child2, opts)

		// When
		var foundTokenFunc *globalTokenFunc
		result := getOptionFromChildren(root, func(_ any, allOpts Options) any {
			if allOpts.tokenFunc != nil {
				foundTokenFunc = &allOpts.tokenFunc
			}
			return foundTokenFunc
		})
		assert.Assert(t, result != nil)
		tokenFunc, ok := result.(*globalTokenFunc)
		assert.Assert(t, ok)
		token, err := (*tokenFunc)(child2)
		assert.NilError(t, err)
		assert.Equal(t, token, "child2-token")
	})

	t.Run("update function can accumulate values", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		child1 := &cobra.Command{Use: "child1"}
		child2 := &cobra.Command{Use: "child2"}
		root.AddCommand(child1)
		root.AddCommand(child2)

		opts1 := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "token1", nil
			},
		}
		opts2 := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "token2", nil
			},
		}

		setCommandOptions(child1, opts1)
		setCommandOptions(child2, opts2)

		// When
		var count int
		result := getOptionFromChildren(root, func(_ any, allOpts Options) any {
			if allOpts.tokenFunc != nil {
				count++
			}
			return &count
		})

		// Then
		assert.Assert(t, result != nil)
		countPtr, ok := result.(*int)
		assert.Assert(t, ok)
		assert.Equal(t, *countPtr, 2)
	})

	t.Run("handles commands with nil context", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		child := &cobra.Command{Use: "child"}
		root.AddCommand(child)
		// child has nil context

		// When
		result := getOptionFromChildren(root, func(currentVal any, _ Options) any {
			return currentVal
		})
		// Then
		assert.Assert(t, result == nil)
	})

	t.Run("handles context without options", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		root.SetContext(context.WithValue(context.Background(), contextKey{}, &contextValue{}))

		// When
		result := getOptionFromChildren(root, func(currentVal any, _ Options) any {
			return currentVal
		})

		// Then
		assert.Assert(t, result == nil)
	})
}

func TestGetGlobalTokenFunc(t *testing.T) {
	t.Run("returns nil when no token func is set", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		result := getGlobalTokenFunc(root)
		// Then
		assert.Assert(t, result == nil)
	})

	t.Run("returns token func when set on root", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		opts := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "root-token", nil
			},
		}
		setCommandOptions(root, opts)

		// When
		tokenFunc := getGlobalTokenFunc(root)

		// Then
		assert.Assert(t, tokenFunc != nil)
		token, err := tokenFunc(root)
		assert.NilError(t, err)
		assert.Equal(t, token, "root-token")
	})

	t.Run("returns token func when set on child", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		child := &cobra.Command{Use: "child"}
		root.AddCommand(child)

		opts := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "child-token", nil
			},
		}
		setCommandOptions(child, opts)

		// When
		tokenFunc := getGlobalTokenFunc(child)

		// Then
		assert.Assert(t, tokenFunc != nil)
		token, err := tokenFunc(child)
		assert.NilError(t, err)
		assert.Equal(t, token, "child-token")
	})

	t.Run("returns nil when token func is nil in options", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		opts := Options{
			tokenFunc: nil,
		}
		setCommandOptions(root, opts)

		// When
		result := getGlobalTokenFunc(root)

		// Then
		assert.Assert(t, result == nil)
	})

	t.Run("warns and returns first token func when multiple are set", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		child1 := &cobra.Command{Use: "child1"}
		child2 := &cobra.Command{Use: "child2"}
		root.AddCommand(child1)
		root.AddCommand(child2)

		opts1 := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "token1", nil
			},
		}
		opts2 := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "token2", nil
			},
		}
		setCommandOptions(child1, opts1)
		setCommandOptions(child2, opts2)

		var output strings.Builder
		child2.SetErr(&output)

		// When
		tokemFunc := getGlobalTokenFunc(child2)

		// Then
		assert.Assert(t, tokemFunc != nil)
		// Should return the first one found (child1's token func)
		token, err := tokemFunc(child2)
		assert.NilError(t, err)
		assert.Equal(t, token, "token1")
		// Should print warning
		assert.Assert(t, strings.Contains(output.String(), "WARNING"))
		assert.Assert(t, strings.Contains(output.String(), "child2"))
	})

	t.Run("returns token func from highest in hierarchy when multiple are set", func(t *testing.T) {
		// Given
		root := &cobra.Command{Use: "root"}
		child1 := &cobra.Command{Use: "child1"}
		child2 := &cobra.Command{Use: "child2"}
		root.AddCommand(child1)
		child1.AddCommand(child2)

		opts1 := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "token1", nil
			},
		}
		opts2 := Options{
			tokenFunc: func(_ *cobra.Command) (string, error) {
				return "token2", nil
			},
		}
		setCommandOptions(root, opts1)
		setCommandOptions(child2, opts2)

		var output strings.Builder
		child2.SetErr(&output)

		// When
		tokenFunc := getGlobalTokenFunc(child2)

		// Then
		assert.Assert(t, tokenFunc != nil)
		// Should return the first one found (root's token func)
		token, err := tokenFunc(child2)
		assert.NilError(t, err)
		assert.Equal(t, token, "token1")
		// Should print warning
		assert.Assert(t, strings.Contains(output.String(), "WARNING"))
	})
}
