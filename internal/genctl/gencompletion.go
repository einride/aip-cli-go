package genctl

import (
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateCompletionFile(gen *protogen.Plugin, rootPackage string) error {
	module := getModule(gen)
	log.Println(rootPackage)
	log.Println(module)
	g := gen.NewGeneratedFile(module+"/completion.go", "")
	cobra := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "github.com/spf13/cobra",
		GoName:       "Command",
	})
	os := g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: "os",
		GoName:       "",
	})
	g.P("package ", rootPackage)
	g.P()
	g.P("var completionCmd = &", cobra, "{")
	g.P(`Use: "completion [bash|zsh|fish|powershell]",`)
	g.P(`Short: "Generate completion script",`)
	g.P(`DisableFlagsInUseLine: true,`)
	g.P(`ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},`)
	g.P(`Args:                  cobra.ExactValidArgs(1),`)
	g.P(`Run: func(cmd *cobra.Command, args []string) {`)
	g.P(`  switch args[0] {`)
	g.P(`  case "bash":`)
	g.P(`    cmd.Root().GenBashCompletion(`, os, `Stdout)`)
	g.P(`  case "zsh":`)
	g.P(`    cmd.Root().GenZshCompletion(`, os, `Stdout)`)
	g.P(`  case "fish":`)
	g.P(`    cmd.Root().GenFishCompletion(`, os, `Stdout, true)`)
	g.P(`  case "powershell":`)
	g.P(`    cmd.Root().GenPowerShellCompletionWithDesc(`, os, `Stdout)`)
	g.P(`  };`)
	g.P(`},`)

	g.P("Long: `To load completions:")
	g.P("Bash:")
	g.P()
	g.P("  $ source <(", "einridectl2", " completion bash)")
	g.P()
	g.P("  # To load completions for each session, execute once:")
	g.P("  # Linux:")
	g.P("  $ ", "einridectl2", " completion bash > /etc/bash_completion.d/", "einridectl2")
	g.P("  # macOS:")
	g.P("  $ ", "einridectl2", " completion bash > /usr/local/etc/bash_completion.d/", "einridectl2")
	g.P()
	g.P("Zsh:")
	g.P()
	g.P("  # If shell completion is not already enabled in your environment,")
	g.P("  # you will need to enable it.  You can execute the following once:")
	g.P()
	g.P(`  $ echo "autoload -U compinit; compinit" >> ~/.zshrc`)
	g.P()
	g.P("  # To load completions for each session, execute once:")
	g.P("  $ ", "einridectl2", ` completion zsh > "${fpath[1]}/_`, "einridectl2")
	g.P()
	g.P("  # You will need to start a new shell for this setup to take effect.")
	g.P()
	g.P("fish:")
	g.P()
	g.P("  $ ", "einridectl2", " completion fish | source")
	g.P()
	g.P("  # To load completions for each session, execute once:")
	g.P("  $ ", "einridectl2", " completion fish > ~/.config/fish/completions/", "einridectl2", ".fish")
	g.P()
	g.P("PowerShell:")
	g.P()
	g.P("  PS> ", "einridectl2", " completion powershell | Out-String | Invoke-Expression")
	g.P()
	g.P("  # To load completions for every new session, run:")
	g.P("  PS> ", "einridectl2", " completion powershell > ", "einridectl2", ".ps1")
	g.P("  # and source this file from your PowerShell profile.")
	g.P("`,")
	g.P("}")
	g.P("func AddCompletion(parent *", cobra, ") {")
	g.P("parent.AddCommand(completionCmd)")
	g.P("}")
	g.P()
	return nil
}
