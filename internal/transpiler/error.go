package transpiler

import (
	"algo-iut/internal/scan"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Error struct {
	src             string
	errStr          string
	s               scan.Scanner
	transpilerStack string
}

func (e *Error) Show(w io.Writer) {
	lines := strings.Split(e.src, "\n")
	line := lines[e.s.Pos().Line-1]

	fmt.Fprintf(w, "Transpiler error: line %v\n", e.s.Pos().Line)
	fmt.Fprintln(w, line)
	fmt.Fprint(w, strings.Repeat(" ", e.s.Pos().Column+1)+"^")
	fmt.Fprintln(w, strings.Repeat("-", len(e.s.Peek())))
	fmt.Fprintln(w, e.errStr)

	fmt.Fprintln(w, "Transpiler stacktrace:")
	fmt.Fprintln(w, e.transpilerStack)

}

func (e *Error) Error() string {
	buf := bytes.Buffer{}
	e.Show(&buf)
	return buf.String()
}
