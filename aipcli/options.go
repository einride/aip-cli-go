package aipcli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type OptFunc func(*Options)

type globalTokenFunc func(cmd *cobra.Command) (string, error)

type Options struct {
	tokenFunc globalTokenFunc
}

type optionsContextKey struct{}

type optionsContextValue struct {
	options Options
}

// WithGlobalTokenFunc allows defining how a bearer token is obtained for all commands which
// perform gRPC calls to the target service.
// If WithGlobalTokenFunc is defined on more than 1 command in the command hierarchy, the highest
// one in the hierarchy will be the one used.
// If no token function is defined, the default token function will be used.
func WithGlobalTokenFunc(tokenFunc func(cmd *cobra.Command) (string, error)) OptFunc {
	return func(o *Options) {
		o.tokenFunc = tokenFunc
	}
}

// getOptionFromChildren walks through the entire tree of commands starting with the given cmd,
// looking for an option of type T in each command's context. If it finds one, it calls the
// given update function, allowing it to decide what to return and eventually returning the final value.
func getOptionFromChildren(cmd *cobra.Command, updateFunc func(currentVal any, allOpts Options) any) any {
	var val any
	var update func(c *cobra.Command)
	update = func(cmd *cobra.Command) {
		ctx := cmd.Context()
		// If we have a context, with an optionsContextValue, update the value with the update function.
		if ctx != nil {
			if v := ctx.Value(optionsContextKey{}); v != nil {
				if typedVal, ok := v.(*optionsContextValue); ok {
					val = updateFunc(val, typedVal.options)
				}
			}
		}
		for _, child := range cmd.Commands() {
			update(child)
		}
	}
	update(cmd)

	return val
}

// setCommandOptions sets the Options in the given command's context.
func setCommandOptions(cmd *cobra.Command, options Options) {
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	cmd.SetContext(context.WithValue(ctx, optionsContextKey{}, &optionsContextValue{
		options: options,
	}))
}

func defaultTokenFunc(cmd *cobra.Command) (string, error) {
	if GetConfig(cmd).CachedIdentityTokenPath != "" {
		tokenFile := GetConfig(cmd).CachedIdentityTokenPath
		identityToken, err := identityTokenFromConfigFile(tokenFile)
		if err != nil {
			return "", fmt.Errorf("failed to read identity token from config file: %v", err)
		}
		return identityToken, nil
	}

	if GetConfig(cmd).GoogleCloudIdentityTokens {
		identityToken, ok := gcloudAuthPrintIdentityToken()
		if !ok {
			return "", fmt.Errorf("failed to print identity token")
		}
		return identityToken, nil
	}

	return "", nil
}

// getGlobalTokenFunc walks down the command hierarchy from the root cmd,
// looking for a global token function in each command's context. If it finds one, it returns it.
// If it doesn't find one, it returns nil.
// Validation is performed to ensure that only one global token function has set.
func getGlobalTokenFunc(cmd *cobra.Command) globalTokenFunc {
	updateFunc := func(currentVal any, opts Options) any {
		if currentVal != nil {
			cmd.PrintErrf(
				//nolint:lll
				`WARNING: command %s attempted to configure a global token function (WithGlobalTokenFunc) when another command has already done so. This is currently not supported and will be ignored.`,
				cmd.Name(),
			)
			return currentVal
		}
		if opts.tokenFunc == nil {
			return nil
		}
		return &opts.tokenFunc
	}
	val := getOptionFromChildren(
		cmd.Root(),
		updateFunc,
	)
	if val == nil {
		return nil
	}
	f, ok := val.(*globalTokenFunc)
	if !ok {
		return nil
	}
	return *f
}
