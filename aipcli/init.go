package aipcli

import "github.com/spf13/cobra"

//nolint: gochecknoinits // need to initialize Cobra's global state
func init() {
	cobra.EnableCommandSorting = false
}
