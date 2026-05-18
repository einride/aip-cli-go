package aipcli

import (
	"context"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Config is the configuration for the protoc-gen-go-aip-cli plugin.
type Config struct {
	// Hosts is a map from host ID to host address.
	Hosts map[string]string
	// DefaultHost is the host ID of the default host.
	DefaultHost string
	// RootPath is where the root file will be located.
	RootPath string
	// Root is the name of the root module package.
	Root string
	// GoogleCloudIdentityTokens indicates if Google Cloud Identity Tokens should be automatically generated.
	GoogleCloudIdentityTokens bool
	// Uses in a predetermined file relative to local config path for the Identity Token.
	CachedIdentityTokenPath string
}

const (
	addressFlag    = "address"
	tokenFlag      = "token"
	insecureFlag   = "insecure"
	verboseFlag    = "verbose"
	forceTraceFlag = "force-trace"
)

type contextKey struct{}

type contextValue struct {
	config               Config
	defaultHostFromProto string
}

func initContext(cmd *cobra.Command, config Config) {
	ctx := cmd.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	cmd.SetContext(context.WithValue(ctx, contextKey{}, &contextValue{
		config: config,
	}))
	cmd.SetOut(os.Stdout)
}

func initPersistentFlags(cmd *cobra.Command) {
	if cmd.PersistentFlags().Lookup(addressFlag) == nil {
		cmd.PersistentFlags().String(addressFlag, "", "Connect to address")
	}
	if cmd.PersistentFlags().Lookup(insecureFlag) == nil {
		cmd.PersistentFlags().Bool(insecureFlag, false, "Make insecure connection (only on localhost)")
	}
	if cmd.PersistentFlags().Lookup(tokenFlag) == nil {
		cmd.PersistentFlags().String(tokenFlag, "", "Authenticate with a bearer token")
	}
	if cmd.PersistentFlags().Lookup(verboseFlag) == nil {
		cmd.PersistentFlags().BoolP(verboseFlag, "v", false, "Enable verbose output")
	}
	if cmd.PersistentFlags().Lookup(forceTraceFlag) == nil {
		cmd.PersistentFlags().BoolP(forceTraceFlag, "", false, "Enable forced Google Cloud Trace header")
	}
	if cmd.PersistentFlags().Lookup("help") == nil {
		cmd.PersistentFlags().BoolP("help", "h", false, "Show help for command")
	}
}

func setConfig(cmd *cobra.Command, config Config) {
	initContext(cmd, config)
	for host, hostAddress := range config.Hosts {
		if cmd.PersistentFlags().Lookup(host) == nil {
			help := "Connect to " + host + " host (" + hostAddress + ")"
			if host == config.DefaultHost {
				help += " [default]"
			}
			cmd.PersistentFlags().Bool(host, false, help)
			_ = cmd.PersistentFlags().SetAnnotation(host, flagHostAnnotation, []string{hostAddress})
		}
	}
}

func setDefaultHostFromProto(cmd *cobra.Command, method protoreflect.MethodDescriptor) {
	if value, ok := cmd.Context().Value(contextKey{}).(*contextValue); ok {
		if defaultHost, ok := proto.GetExtension(
			method.Parent().Options(),
			annotations.E_DefaultHost,
		).(string); ok {
			value.defaultHostFromProto = defaultHost
		}
	}
}

func defaultHostFromProto(cmd *cobra.Command) (string, bool) {
	result := cmd.Context().Value(contextKey{}).(*contextValue).defaultHostFromProto
	return result, result != ""
}

func getContext(cmd *cobra.Command) (context.Context, bool) {
	ctx := cmd.Context()
	if ctx != nil {
		return ctx, true
	}
	parent := cmd.Parent()
	for parent != nil {
		ctx = parent.Context()
		if ctx != nil {
			return ctx, true
		}
		parent = parent.Parent()
	}
	return nil, false
}

func GetConfig(cmd *cobra.Command) Config {
	ctx, ok := getContext(cmd)
	if !ok {
		return Config{}
	}
	if value, ok := ctx.Value(contextKey{}).(*contextValue); ok {
		return value.config
	}
	return Config{}
}

// lookupChangedPersistentFlag walks cmd and its ancestors looking for a persistent flag
// with the given name that the user explicitly set (Changed == true). Using
// PersistentFlags() rather than Flags() means local proto-field flags (registered via
// cmd.Flags()) are never returned, so starting from cmd itself is safe even when a proto
// field shares a name with a connection flag.
func lookupChangedPersistentFlag(cmd *cobra.Command, name string) *pflag.Flag {
	for c := cmd; c != nil; c = c.Parent() {
		if f := c.PersistentFlags().Lookup(name); f != nil && f.Changed {
			return f
		}
	}
	return nil
}

func getAddress(cmd *cobra.Command) (string, bool) {
	if f := lookupChangedPersistentFlag(cmd, addressFlag); f != nil {
		return f.Value.String(), true
	}
	for host, hostAddress := range GetConfig(cmd).Hosts {
		if useHost, err := cmd.Flags().GetBool(host); err == nil && useHost {
			return hostAddress, true
		}
	}
	if defaultHostFromConfig := GetConfig(cmd).DefaultHost; defaultHostFromConfig != "" {
		for host, hostAddress := range GetConfig(cmd).Hosts {
			if host == defaultHostFromConfig {
				return hostAddress, true
			}
		}
	}
	return defaultHostFromProto(cmd)
}

func getToken(cmd *cobra.Command) (string, error) {
	if flagToken, err := cmd.Flags().GetString(tokenFlag); err == nil && flagToken != "" {
		return flagToken, nil
	}

	tokenFunc := getGlobalTokenFunc(cmd)
	if tokenFunc != nil {
		return tokenFunc(cmd)
	}

	token, err := defaultTokenFunc(cmd)
	if err != nil {
		return "", err
	}
	return token, nil
}

func isInsecure(cmd *cobra.Command) bool {
	if f := lookupChangedPersistentFlag(cmd, insecureFlag); f != nil {
		v, _ := strconv.ParseBool(f.Value.String())
		return v
	}
	return false
}

func IsVerbose(cmd *cobra.Command) bool {
	if f := lookupChangedPersistentFlag(cmd, verboseFlag); f != nil {
		v, _ := strconv.ParseBool(f.Value.String())
		return v
	}
	return false
}

func isForceTrace(cmd *cobra.Command) bool {
	if f := lookupChangedPersistentFlag(cmd, forceTraceFlag); f != nil {
		v, _ := strconv.ParseBool(f.Value.String())
		return v
	}
	return false
}
