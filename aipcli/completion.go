package aipcli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"go.einride.tech/aip/resourcename"
)

type CompletionFunc func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)

func ResourceNameCompletionFunc(patterns ...string) CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		result := make([]string, 0, len(patterns))
		for _, pattern := range patterns {
			if completion, ok := CompleteResourceName(toComplete, pattern); ok {
				result = append(result, fmt.Sprintf("%s\t%s", completion, pattern))
			}
		}
		return result, cobra.ShellCompDirectiveNoSpace
	}
}

func ResourceNameListCompletionFunc(patterns ...string) CompletionFunc {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		toCompleteElements := strings.Split(toComplete, ",")
		lastToCompleteElement := toCompleteElements[len(toCompleteElements)-1]
		result := make([]string, 0, len(patterns))
		for _, pattern := range patterns {
			if elementCompletion, ok := CompleteResourceName(lastToCompleteElement, pattern); ok {
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
		return result, cobra.ShellCompDirectiveNoSpace
	}
}

func CompleteResourceName(toComplete, pattern string) (string, bool) {
	toCompleteSegments := strings.Split(toComplete, "/")
	patternSegments := strings.Split(pattern, "/")
	if len(toCompleteSegments) > len(patternSegments) {
		return "", false
	}
	var result strings.Builder
	result.Grow(len(pattern))
	for i, toCompleteSegment := range toCompleteSegments {
		patternSegment := patternSegments[i]
		if resourcename.Segment(patternSegment).IsVariable() {
			result.WriteString(toCompleteSegment)
			if i < len(toCompleteSegments)-1 {
				result.WriteByte('/')
			}
			continue
		}
		if toCompleteSegment == patternSegment {
			result.WriteString(patternSegment)
			if i < len(toCompleteSegments)-1 || i < len(patternSegments)-1 {
				result.WriteByte('/')
			}
			continue
		}
		if i < len(toCompleteSegments)-1 {
			return "", false
		}
		if !strings.HasPrefix(patternSegment, toCompleteSegment) {
			return "", false
		}
		result.WriteString(patternSegment)
		if i < len(patternSegments)-1 {
			result.WriteByte('/')
		}
	}
	return result.String(), result.String() != ""
}
