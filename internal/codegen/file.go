package codegen

import (
	"bytes"
	"fmt"
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

func (f *File) Content() []byte {
	return f.buf.Bytes()
}
