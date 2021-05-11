package codegen

import (
	"bytes"
	"path"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Imports struct {
	// path -> identifier
	usedImports     map[string]string
	usedIdentifiers map[string]struct{}
}

func (i *Imports) Import(path string) string {
	if i.usedImports == nil {
		i.usedImports = make(map[string]string)
	}
	if i.usedIdentifiers == nil {
		i.usedIdentifiers = make(map[string]struct{})
	}
	if identifier, ok := i.usedImports[path]; ok {
		return identifier
	}
	identifier := importPathToAssumedName(path)
	for {
		if _, ok := i.usedIdentifiers[identifier]; !ok {
			break
		}
		identifier += "x"
	}
	i.usedImports[path] = identifier
	return identifier
}

func (i Imports) Bytes() []byte {
	if len(i.usedImports) == 0 {
		return nil
	}
	type pkg struct {
		importPath, packageName string
	}
	stdPkgs := make([]pkg, 0, len(i.usedImports))
	nonStdPkgs := make([]pkg, 0, len(i.usedImports))
	for importPath, packageName := range i.usedImports {
		if strings.Contains(importPath, ".") && packageName == path.Base(importPath) {
			nonStdPkgs = append(nonStdPkgs, pkg{importPath: importPath, packageName: packageName})
		} else {
			stdPkgs = append(stdPkgs, pkg{importPath: importPath, packageName: packageName})
		}
	}
	sort.Slice(stdPkgs, func(i, j int) bool {
		return stdPkgs[i].importPath < stdPkgs[j].importPath
	})
	sort.Slice(nonStdPkgs, func(i, j int) bool {
		return nonStdPkgs[i].importPath < nonStdPkgs[j].importPath
	})
	var b bytes.Buffer
	_ = b.WriteByte('\n')
	_, _ = b.WriteString("import (")
	for _, stdPkg := range stdPkgs {
		_, _ = b.WriteString(strconv.Quote(stdPkg.importPath))
		_ = b.WriteByte('\n')
	}
	if len(nonStdPkgs) > 0 {
		_ = b.WriteByte('\n')
	}
	for _, nonStdPkg := range nonStdPkgs {
		if nonStdPkg.packageName == path.Base(nonStdPkg.importPath) {
			_, _ = b.WriteString(strconv.Quote(nonStdPkg.importPath))
			_ = b.WriteByte('\n')
		} else {
			_, _ = b.WriteString(nonStdPkg.packageName)
			_ = b.WriteByte(' ')
			_, _ = b.WriteString(strconv.Quote(nonStdPkg.importPath))
			_ = b.WriteByte('\n')
		}
	}
	_ = b.WriteByte(')')
	return b.Bytes()
}

// importPathToAssumedName is copy-pasted from golang.org/x/tools.
func importPathToAssumedName(importPath string) string {
	base := path.Base(importPath)
	if strings.HasPrefix(base, "v") {
		if _, err := strconv.Atoi(base[1:]); err == nil {
			dir := path.Dir(importPath)
			if dir != "." {
				base = path.Base(dir)
			}
		}
	}
	base = strings.TrimPrefix(base, "go-")
	if i := strings.IndexFunc(base, func(r rune) bool {
		return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || '0' <= r && r <= '9' || r == '_' ||
			r >= utf8.RuneSelf && (unicode.IsLetter(r) || unicode.IsDigit(r)))
	}); i >= 0 {
		base = base[:i]
	}
	return base
}
