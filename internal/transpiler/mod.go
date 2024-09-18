package transpiler

import (
	"fmt"
	"io"
	"text/scanner"
)

func DoRoot(s *scanner.Scanner, output io.WriteCloser) {
	// first scan
	s.Scan()

	for s.Peek() != scanner.EOF {
		switch s.TokenText() {
		case "fonction":
			s.Scan()
			doFunction(s, output)
		case "procedure":
			s.Scan()
			doProcedure(s, output)
		default:
			panic(fmt.Sprintf("unexpected token '%s'", s.TokenText()))
		}
	}
}
