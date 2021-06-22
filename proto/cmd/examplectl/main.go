package main

import (
	context "context"
	fmt "fmt"
	cli "go.einride.tech/protoc-gen-go-cli/cli"
	os "os"
)

func main() {
	ctx := context.Background()
	cmd := NewRootCommand()
	config := NewConfig()
	config.Runtime.AddToFlagSet(cmd.PersistentFlags(), config.Compiler)
	ctx = cli.WithConfig(ctx, config)
	if err := cmd.ExecuteContext(ctx); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
