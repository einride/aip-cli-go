package aipcli

import (
	"context"
	"os/exec"
	"strings"
)

type Config struct {
	Compiler CompilerConfig
	Runtime  RuntimeConfig
}

func (c *Config) GetAddress() (string, bool) {
	if c.Runtime.Address != "" {
		return c.Runtime.Address, true
	}
	for host, ok := range c.Runtime.Host {
		if !*ok {
			continue
		}
		return c.Compiler.Hosts[host], true
	}
	if defaultAddress, ok := c.Compiler.DefaultAddress(); ok {
		return defaultAddress, true
	}
	if c.Runtime.DefaultHost != "" {
		return c.Runtime.DefaultHost, true
	}
	return "", false
}

func (c *Config) GetToken() (string, bool) {
	if c.Runtime.Token != "" {
		return c.Runtime.Token, true
	}
	if c.Compiler.GoogleCloudIdentityTokens {
		return getGoogleCloudIdentityToken()
	}
	return "", false
}

func getGoogleCloudIdentityToken() (string, bool) {
	if _, err := exec.LookPath("gcloud"); err != nil {
		return "", false
	}
	var stdout strings.Builder
	cmd := exec.Command("gcloud", "auth", "print-identity-token")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", false
	}
	return strings.TrimSpace(stdout.String()), true
}

func ConfigFromContext(ctx context.Context) *Config {
	config, ok := ctx.Value(configContextKey{}).(*Config)
	if !ok {
		panic("aipcli.ConfigFromContext was called with a context without a Config")
	}
	return config
}

func WithConfig(ctx context.Context, config *Config) context.Context {
	return context.WithValue(ctx, configContextKey{}, config)
}

func SetDefaultHost(ctx context.Context, host string) {
	ConfigFromContext(ctx).Runtime.DefaultHost = host
}

type configContextKey struct{}
