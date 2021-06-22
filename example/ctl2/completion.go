package ctl2

import (
	os "os"

	cobra "github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
	Long: `To load completions:
Bash:

  $ source <(einridectl2 completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ einridectl2 completion bash > /etc/bash_completion.d/einridectl2
  # macOS:
  $ einridectl2 completion bash > /usr/local/etc/bash_completion.d/einridectl2

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ einridectl2 completion zsh > "${fpath[1]}/_einridectl2

  # You will need to start a new shell for this setup to take effect.

Fish:

  $ einridectl2 completion fish | source

  # To load completions for each session, execute once:
  $ einridectl2 completion fish > ~/.config/fish/completions/einridectl2.fish

PowerShell:

  PS> einridectl2 completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> einridectl2 completion powershell > einridectl2.ps1
  # and source this file from your PowerShell profile.
`,
}

func AddCompletion(parent *cobra.Command) {
	parent.AddCommand(completionCmd)
}
