package aipcli

import (
	"fmt"

	"github.com/spf13/pflag"
)

type RuntimeConfig struct {
	Verbose     bool
	Insecure    bool
	Token       string
	Address     string
	Host        map[string]*bool
	DefaultHost string
}

func (c *RuntimeConfig) AddToFlagSet(flags *pflag.FlagSet, compiler CompilerConfig) {
	flags.BoolVarP(&c.Verbose, "verbose", "v", false, "print verbose output")
	flags.BoolVar(&c.Insecure, "insecure", false, "make insecure client connection (only on localhost)")
	flags.StringVar(&c.Address, "address", "", "address to connect to")
	flags.StringVar(&c.Token, "token", "", "bearer token used by client")
	c.Host = map[string]*bool{}
	for host, address := range compiler.Hosts {
		if flag := flags.Lookup(host); flag == nil {
			var hostFlag bool
			flags.BoolVar(&hostFlag, host, false, fmt.Sprintf("connect to %s host (%s)", host, address))
			c.Host[host] = &hostFlag
		}
	}
}
