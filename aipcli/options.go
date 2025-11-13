package aipcli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

type OptFunc func(*Options)

type Options struct {
	tokenFunc func(cmd *cobra.Command) (string, error)
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

type optionsContextKey struct{}

type optionsContextValue struct {
	options Options
}

func collectChildrenOptions(cmd *cobra.Command) Options {
	// Walk through all command children and their children, merging Options.
	// If a second tokenFunc is encountered, error out.
	var tokenFunc func(cmd *cobra.Command) (string, error)

	// Function will return (Options, error). Since the original signature is void, we use panic for error.
	var collect func(cmd *cobra.Command)
	collect = func(cmd *cobra.Command) {
		opt := getContextOptions(cmd.Context())

		if opt.tokenFunc != nil {
			if tokenFunc != nil {
				cmd.PrintErrf(
					//nolint:lll
					`WARNING: command %s attempted to configure a global token function (WithGlobalTokenFunc) when another command has already done so. This is currently not supported and will be ignored.`,
					cmd.Name(),
				)
			} else {
				tokenFunc = opt.tokenFunc
			}
		}

		for _, child := range cmd.Commands() {
			collect(child)
		}
	}

	collect(cmd)

	return Options{
		tokenFunc: tokenFunc,
	}
}

func getContextOptions(ctx context.Context) Options {
	if ctx == nil {
		return Options{}
	}
	ctxValue := ctx.Value(optionsContextKey{})
	if ctxValue == nil {
		return Options{}
	}
	value, ok := ctxValue.(*optionsContextValue)
	if !ok {
		return Options{}
	}
	return value.options
}

func setCommandOptions(cmd *cobra.Command, options Options) {
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	cmd.SetContext(context.WithValue(ctx, optionsContextKey{}, &optionsContextValue{
		options: options,
	}))
}

func getRootOptions(cmd *cobra.Command) Options {
	return getContextOptions(cmd.Root().Context())
}

func getOptions(cmd *cobra.Command) Options {
	return getRootOptions(cmd)
}
