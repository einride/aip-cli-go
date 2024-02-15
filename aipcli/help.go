package aipcli

import (
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func helpFunc(cmd *cobra.Command, _ []string) {
	_ = usageFunc(cmd)
}

func usageFunc(cmd *cobra.Command) error {
	out := cmd.ErrOrStderr()
	defer cmd.SetErr(out)
	tw := tabwriter.NewWriter(out, 2, 0, 2, ' ', 0)
	cmd.SetErr(tw)
	if cmd.Short != "" {
		cmd.PrintErrln()
		if cmd.Long != "" {
			cmd.PrintErrln(cmd.Long)
		} else {
			cmd.PrintErrln(cmd.Short)
		}
	}
	cmd.PrintErrln()
	cmd.PrintErrln("USAGE")
	cmd.PrintErrln(" ", cmd.Use, "<command>")
	if methods := getCommands(cmd, commandHasAnnotation(methodAnnotation)); len(methods) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("METHOD COMMANDS")
		printCommands(cmd, methods)
	}
	if services := getCommands(cmd, commandHasAnnotation(serviceAnnotation)); len(services) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("SERVICE COMMANDS")
		printCommands(cmd, services)
	}
	if modules := getCommands(cmd, commandHasAnnotation(moduleAnnotation)); len(modules) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("MODULE COMMANDS")
		printCommands(cmd, modules)
	}
	if otherCommands := getCommands(cmd, func(cmd *cobra.Command) bool {
		_, isModule := cmd.Annotations[moduleAnnotation]
		_, isService := cmd.Annotations[serviceAnnotation]
		_, isMethod := cmd.Annotations[methodAnnotation]
		return !isModule && !isService && !isMethod
	}); len(otherCommands) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("OTHER COMMANDS")
		printCommands(cmd, otherCommands)
	}
	if argumentFlags := getFlags(cmd, isArgumentFlag); len(argumentFlags) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("ARGUMENT FLAGS")
		printFlags(cmd, argumentFlags)
	}
	if hostFlags := getFlags(cmd, isHostFlag); len(hostFlags) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("HOST FLAGS")
		printFlags(cmd, hostFlags)
	}
	if connectionFlags := getFlags(cmd, isConnectionFlag); len(connectionFlags) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("CONNECTION FLAGS")
		printFlags(cmd, connectionFlags)
	}
	if otherFlags := getFlags(cmd, func(_ *cobra.Command, flag *pflag.Flag) bool {
		return !isConnectionFlag(cmd, flag) && !isArgumentFlag(cmd, flag) && !isHostFlag(cmd, flag)
	}); len(otherFlags) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("OTHER FLAGS")
		printFlags(cmd, otherFlags)
	}
	return tw.Flush()
}

func printCommands(cmd *cobra.Command, commands []*cobra.Command) {
	for _, command := range commands {
		cmd.PrintErrln("  " + command.Name() + "\t" + command.Short)
	}
}

func printFlags(cmd *cobra.Command, flags []*pflag.Flag) {
	var hasShorthand bool
	for _, flag := range flags {
		if flag.Shorthand != "" {
			hasShorthand = true
			break
		}
	}
	var maxFieldBehaviors int
	for _, flag := range flags {
		if curr := len(flag.Annotations[fieldBehaviorsAnnotation]); curr > maxFieldBehaviors {
			maxFieldBehaviors = curr
		}
	}
	for _, flag := range flags {
		if flag.Hidden {
			continue
		}
		var line strings.Builder
		_, _ = line.WriteString("  ")
		if hasShorthand {
			if flag.Shorthand == "" {
				_, _ = line.WriteString("  ")
			} else {
				_, _ = line.WriteString("-" + flag.Shorthand)
			}
			_, _ = line.WriteString("  ")
		}
		_, _ = line.WriteString("--" + flag.Name)
		_, _ = line.WriteString("\t")
		if fieldBehaviors := flag.Annotations[fieldBehaviorsAnnotation]; len(fieldBehaviors) > 0 {
			line.WriteString("[")
			for _, fieldBehavior := range fieldBehaviors {
				line.WriteString(fieldBehavior[:1])
			}
			line.WriteString("]")
		}
		_, _ = line.WriteString("\t")
		_, _ = line.WriteString(flag.Value.Type())
		_, _ = line.WriteString("\t")
		_, _ = line.WriteString(flag.Usage)
		if flag.DefValue != "" && flag.DefValue != "false" {
			_, _ = line.WriteString(" (" + flag.DefValue + ")")
		}
		cmd.PrintErrln(line.String())
	}
}

func isConnectionFlag(_ *cobra.Command, flag *pflag.Flag) bool {
	switch flag.Name {
	case addressFlag, insecureFlag, tokenFlag, forceTraceFlag:
		return true
	}
	return false
}

func isHostFlag(cmd *cobra.Command, flag *pflag.Flag) bool {
	for host := range GetConfig(cmd).Hosts {
		if flag.Name == host {
			return true
		}
	}
	return false
}

func isArgumentFlag(_ *cobra.Command, flag *pflag.Flag) bool {
	_, ok := flag.Annotations[flagArgumentAnnotation]
	return ok
}

func commandHasAnnotation(annotation string) func(command *cobra.Command) bool {
	return func(command *cobra.Command) bool {
		_, ok := command.Annotations[annotation]
		return ok
	}
}

func getCommands(cmd *cobra.Command, fn func(*cobra.Command) bool) []*cobra.Command {
	cmds := cmd.Commands()
	result := make([]*cobra.Command, 0, len(cmds))
	for _, subCmd := range cmds {
		if fn(subCmd) {
			result = append(result, subCmd)
		}
	}
	return result
}

func getFlags(cmd *cobra.Command, fn func(*cobra.Command, *pflag.Flag) bool) []*pflag.Flag {
	flags := cmd.Flags()
	flags.SortFlags = false
	result := make([]*pflag.Flag, 0, flags.NFlag())
	flags.VisitAll(func(flag *pflag.Flag) {
		if fn(cmd, flag) {
			result = append(result, flag)
		}
	})
	return result
}
