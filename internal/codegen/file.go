package codegen

import (
	"bytes"
	"fmt"
	"go/format"
)

type File struct {
	buf bytes.Buffer
}

func (f *File) P(v ...interface{}) {
	for _, x := range v {
		_, _ = fmt.Fprint(&f.buf, x)
	}
	_, _ = fmt.Fprintln(&f.buf)
}

func (f *File) Content() ([]byte, error) {
	return format.Source(f.buf.Bytes())
}
