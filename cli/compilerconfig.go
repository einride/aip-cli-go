package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type CompilerConfig struct {
	Hosts                     map[string]string
	DefaultHost               string
	Root                      string
	GoogleCloudIdentityTokens bool
	Main                      bool
}

func (c *CompilerConfig) DefaultAddress() (string, bool) {
	address, ok := c.Hosts[c.DefaultHost]
	return address, ok
}

func (c *CompilerConfig) AddToFlagSet(flags *pflag.FlagSet) {
	c.Hosts = make(map[string]string)
	flags.Var(stringStringMap{value: c.Hosts}, "hosts", "mapping from alias to host")
	flags.StringVar(&c.DefaultHost, "default_host", "", "default host override")
	flags.StringVar(&c.Root, "root", "", "root command")
	flags.BoolVar(&c.Main, "main", false, "write default root command main.go")
	flags.BoolVar(&c.GoogleCloudIdentityTokens, "gcloud_identity_tokens", false, "use gcloud to print identity tokens")
}

func (c *CompilerConfig) Validate() error {
	if c.DefaultHost != "" && c.Hosts[c.DefaultHost] == "" {
		hosts := make([]string, 0, len(c.Hosts))
		for host := range c.Hosts {
			hosts = append(hosts, host)
		}
		return fmt.Errorf("default_host (%s) must be one of hosts (%s)", c.DefaultHost, strings.Join(hosts, ","))
	}
	if c.Main && c.Root == "" {
		return fmt.Errorf("main option requires root option")
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
