package plugin

import (
	"github.com/einride/ctl/internal/codegen"
)

func GenerateCompletionFile(f *codegen.File) {
	cobra := f.Import("github.com/spf13/cobra")
	os := f.Import("os")
	f.Pf("var completionCmd = &%s.Command{", cobra)
	f.Pf(`Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(%s.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(%s.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(%s.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(%s.Stdout)
		}
	},`, os, os, os, os)

	f.P("Long: `To load completions:")
	f.Pf(`Bash:

  $ source <(%s completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ %s completion bash > /etc/bash_completion.d/%s
  # macOS:
  $ %s completion bash > /usr/local/etc/bash_completion.d/%s

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ %s completion zsh > "${fpath[1]}/_%s"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ %s completion fish | source

  # To load completions for each session, execute once:
  $ %s completion fish > ~/.config/fish/completions/%s.fish

PowerShell:

  PS> %s completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> %s completion powershell > %s.ps1
  # and source this file from your PowerShell profile.`,
		programName,
		programName,
		programName,
		programName,
		programName,
		programName,
		programName,
		programName,
		programName,
		programName,
		programName,
		programName,
		programName,
	)

	f.P("`,")
	f.P("}")

	f.P("func init() {")
	f.P("rootCmd.AddCommand(completionCmd)")
	f.P("}")
}
