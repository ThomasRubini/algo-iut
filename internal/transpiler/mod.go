package transpiler

import (
	"fmt"
	"io"
	"text/scanner"
)

func DoRoot(s *scanner.Scanner, output io.WriteCloser) {
	for s.Scan() != scanner.EOF {
		switch s.TokenText() {
		case "fonction":
			doFunction(s, output)
		case "procedure":
			doProcedure(s, output)
		default:
			panic(fmt.Sprintf("unexpected token '%s'", s.TokenText()))
		}
	}
}
