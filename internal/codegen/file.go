package codegen

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"
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

func (f *File) Pf(format string, a ...interface{}) {
	f.P(fmt.Sprintf(format, a))
}

func (f *File) Pfq(format string, a ...string) {
	ss := make([]string, 0, len(a))
	for _, s := range a {
		ss = append(ss, strconv.Quote(s))
	}
	f.Pf(format, ss)
}

func (f *File) Content() []byte {
	bs := f.buf.Bytes()
	// try to format
	formatted, err := format.Source(bs)
	if err == nil {
		return formatted
	}
	return bs
}
