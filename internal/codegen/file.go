package codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"
)

type File struct {
	pkg  string
	imps Imports
	buf  bytes.Buffer
}

func NewFile(pkg string) *File {
	return &File{
		pkg: pkg,
	}
}

func (f *File) P(v ...interface{}) {
	for _, x := range v {
		_, _ = fmt.Fprint(&f.buf, x)
	}
	_, _ = fmt.Fprintln(&f.buf)
}

func (f *File) Pf(format string, a ...interface{}) {
	f.P(fmt.Sprintf(format, a...))
}

func (f *File) Pfq(format string, a ...string) {
	ss := make([]interface{}, 0, len(a))
	for _, s := range a {
		ss = append(ss, strconv.Quote(s))
	}
	f.Pf(format, ss...)
}

func (f *File) Import(path string) string {
	return f.imps.Import(path)
}

func (f *File) Content() []byte {
	bs := f.content()
	// try to format
	formatted, err := format.Source(bs)
	if err == nil {
		return formatted
	}
	return bs
}

func (f *File) content() []byte {
	var b bytes.Buffer
	_, _ = fmt.Fprintln(&b, "package", f.pkg)
	_, _ = fmt.Fprintln(&b, "// Code generated. DO NOT EDIT.")
	b.Write(f.imps.Bytes())
	b.Write(f.buf.Bytes())
	return b.Bytes()
}
