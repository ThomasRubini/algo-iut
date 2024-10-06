// this package implement the output interface of our transpiler
package langoutput

import (
	"fmt"
	"io"
)

type T interface {
	Write(s string)
	Writef(format string, a ...any)
	Close()
}

type writeCloserOutput struct {
	wc io.WriteCloser
}

func (o *writeCloserOutput) Write(s string) {
	o.wc.Write([]byte(s))
}

func (o *writeCloserOutput) Writef(format string, a ...any) {
	o.Write(fmt.Sprintf(format, a...))
}

func (o *writeCloserOutput) Close() {
	o.wc.Close()
}

func NewWriteCloser(wc io.WriteCloser) T {
	return &writeCloserOutput{
		wc: wc,
	}
}
