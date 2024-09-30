package nopwritecloser

import "io"

type wrapper struct {
	io.Writer
}

func (wc *wrapper) Close() error {
	return nil
}

func New(w io.Writer) io.WriteCloser {
	return &wrapper{
		Writer: w,
	}
}
