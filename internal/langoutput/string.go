package langoutput

import (
	"algo-iut/internal/utils/nopwritecloser"
	"bytes"
)

type WriteCloserString struct {
	T
	buf bytes.Buffer
}

func NewString() WriteCloserString {
	buf := bytes.Buffer{}
	return WriteCloserString{
		T:   NewWriteCloser(nopwritecloser.New(&buf)),
		buf: buf,
	}
}

func (o *WriteCloserString) String() string {
	return o.buf.String()
}
