package protoshell

import (
	"strings"

	"go.einride.tech/aip/resourcename"
)

// CompleteResourceName returns a value to complete for the provided resource name pattern.
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
