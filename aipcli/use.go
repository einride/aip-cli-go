package aipcli

import (
	"bufio"
	"bytes"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func deduplicateServiceCommandUses(cmds []*cobra.Command) {
	sorted := make([]*cobra.Command, 0, len(cmds))
	sorted = append(sorted, cmds...)
	sort.Slice(sorted, func(i, j int) bool {
		// sort descending to get shorter use for higher version
		return sorted[i].Annotations[serviceAnnotation] > sorted[j].Annotations[serviceAnnotation]
	})
	uses := make(map[string]struct{}, len(cmds))
	for _, cmd := range sorted {
		serviceName := protoreflect.FullName(cmd.Annotations[serviceAnnotation])
		if serviceName == "" {
			uses[cmd.Use] = struct{}{}
			continue
		}
		var i int
		for {
			use, ok := qualifiedServiceUse(serviceName, i)
			if !ok {
				break
			}
			_, isUsed := uses[use]
			if !isUsed {
				uses[use] = struct{}{}
				cmd.Use = use
				break
			}
			i++
		}
	}
}

func serviceUse(service protoreflect.FullName) string {
	result := string(service.Name())
	result = strings.TrimSuffix(result, "Service")
	result = strcase.KebabCase(result)
	return result
}

func qualifiedServiceUse(service protoreflect.FullName, n int) (string, bool) {
	trim := service.Parent()
	for i := 0; i < n; i++ {
		if trim == "" {
			return "", false
		}
		trim = trim.Parent()
	}
	result := strings.TrimPrefix(string(service), string(trim)+".")
	result = strings.TrimSuffix(result, "Service")
	result = strings.Join(reverse(strings.Split(result, ".")), "-")
	result = strcase.KebabCase(result)
	return result, true
}

func reverse(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}

func methodUse(method protoreflect.MethodDescriptor) string {
	result := string(method.Name())
	result = strcase.KebabCase(result)
	return result
}

func trimLongComment(s string) string {
	var result strings.Builder
	sc := bufio.NewScanner(strings.NewReader(s))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		_, _ = result.Write(bytes.TrimSpace(sc.Bytes()))
		_ = result.WriteByte('\n')
	}
	return strings.TrimSpace(result.String())
}

func trimFieldComment(comment string) string {
	result := comment
	// Clean up comment line breaks and whitespace.
	result = strings.ReplaceAll(result, "//", "")
	result = strings.ReplaceAll(result, "\n", " ")
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "  ", " ")
	result = strings.ReplaceAll(result, "  ", " ")
	// Cut out first sentence.
	if i := strings.IndexByte(result, '.'); i != -1 {
		result = result[:i]
	}
	// Trim manually documented field behavior.
	result = strings.TrimPrefix(result, "REQUIRED: ")
	result = strings.TrimPrefix(result, "Required: ")
	result = strings.TrimPrefix(result, "The")
	result = strings.TrimPrefix(result, "the")
	result = strings.TrimSpace(result)
	result = initialUpperCase(result)
	return result
}

func initialUpperCase(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if size == utf8.RuneError {
		return s
	}
	return string(unicode.ToUpper(r)) + s[size:]
}
