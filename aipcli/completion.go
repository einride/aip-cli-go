package aipcli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"go.einride.tech/aip-cli/internal/protoshell"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type CompletionFunc func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective)

func getResourceNameActiveHelp(comment string, patterns ...string) string {
	result := trimFieldComment(comment)
	filteredPatterns := make([]string, 0, len(patterns))
	for _, pattern := range patterns {
		if !strings.Contains(result, pattern) {
			filteredPatterns = append(filteredPatterns, pattern)
		}
	}
	if len(filteredPatterns) > 0 {
		result += " (" + strings.Join(patterns, " or ") + ")"
	}
	return result
}

func fieldCompletionFunc(comment string) CompletionFunc {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return cobra.AppendActiveHelp(nil, trimFieldComment(comment)),
			cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	}
}

func enumFieldCompletionFunc(comment string, values protoreflect.EnumValueDescriptors) CompletionFunc {
	return func(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return cobra.AppendActiveHelp(protoshell.CompleteEnumValue(toComplete, values), trimFieldComment(comment)),
			cobra.ShellCompDirectiveNoFileComp
	}
}

func timestampCompletionFunc(comment string) CompletionFunc {
	return func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		activeHelp := trimFieldComment(comment)
		activeHelp += " [tip: prefix with = to evaluate a CEL expression, e.g. ='now()-duration(\"2h\")']"
		return cobra.AppendActiveHelp(nil, activeHelp),
			cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	}
}

func resourceNameCompletionFunc(comment string, patterns ...string) CompletionFunc {
	return func(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		result := make([]string, 0, len(patterns))
		for _, pattern := range patterns {
			if completion, ok := protoshell.CompleteResourceName(toComplete, pattern); ok {
				result = append(result, fmt.Sprintf("%s\t%s", completion, pattern))
			}
		}
		result = cobra.AppendActiveHelp(result, getResourceNameActiveHelp(comment, patterns...))
		return result, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	}
}

func resourceNameListCompletionFunc(comment string, patterns ...string) CompletionFunc {
	return func(_ *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		toCompleteElements := strings.Split(toComplete, ",")
		lastToCompleteElement := toCompleteElements[len(toCompleteElements)-1]
		result := make([]string, 0, len(patterns))
		for _, pattern := range patterns {
			if elementCompletion, ok := protoshell.CompleteResourceName(lastToCompleteElement, pattern); ok {
				var completion string
				if len(toCompleteElements) > 1 {
					completion = strings.Join(
						append(toCompleteElements[:len(toCompleteElements)-1], elementCompletion),
						",",
					)
				} else {
					completion = elementCompletion
				}
				result = append(result, fmt.Sprintf("%s\t%s", completion, pattern))
			}
		}
		result = cobra.AppendActiveHelp(result, getResourceNameActiveHelp(comment, patterns...))
		return result, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	}
}
