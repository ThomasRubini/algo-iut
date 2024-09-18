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
		tok := s.TokenText()
		s.Scan()
		switch tok {
		case "fonction":
			doFunction(s, output)
		case "procedure":
			doProcedure(s, output)
		default:
			panic(fmt.Sprintf("unexpected token '%s'", tok))
		}
	}
}
