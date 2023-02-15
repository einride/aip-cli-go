package gengoaipcli

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"go.einride.tech/aip-cli/aipcli"
)

// Config for the protoc-gen-go-aip-cli plugin.
type Config aipcli.Config

// AddToFlagSet adds the config to a pflag.FlagSet.
func (c *Config) AddToFlagSet(flags *pflag.FlagSet) {
	c.Hosts = make(map[string]string)
	flags.Var(stringStringMap{value: c.Hosts}, "hosts", "mapping from alias to host")
	flags.StringVar(&c.DefaultHost, "default_host", "", "default host override")
	flags.StringVar(&c.Root, "root", "", "root command")
	flags.BoolVar(&c.GoogleCloudIdentityTokens, "gcloud_identity_tokens", false, "use gcloud to print identity tokens")
	flags.StringVar(
		&c.CachedIdentityTokenPath,
		"cached_identity_token_path",
		"",
		"a file relative to the local config folder that provides an IdentityToken used by the CLI",
	)
}

// Validate the config.
func (c *Config) Validate() error {
	if c.DefaultHost != "" && c.Hosts[c.DefaultHost] == "" {
		hosts := make([]string, 0, len(c.Hosts))
		for host := range c.Hosts {
			hosts = append(hosts, host)
		}
		return fmt.Errorf("default_host (%s) must be one of hosts (%s)", c.DefaultHost, strings.Join(hosts, ","))
	}

	if c.GoogleCloudIdentityTokens && c.CachedIdentityTokenPath != "" {
		return fmt.Errorf("gcloud_identity_tokens and cached_identity_token_path are mutually exclusive flags")
	}

	return nil
}

type stringStringMap struct {
	value map[string]string
}

func (m stringStringMap) Type() string {
	return "map<string,string>"
}

func (m stringStringMap) String() string {
	return fmt.Sprintf("%v", m.value)
}

func (m stringStringMap) Set(s string) error {
	if m.value == nil {
		m.value = map[string]string{}
	}
	keyValue := strings.SplitN(s, "=", 2)
	if len(keyValue) != 2 {
		return fmt.Errorf("invalid pair: %s", s)
	}
	m.value[keyValue[0]] = keyValue[1]
	return nil
}
