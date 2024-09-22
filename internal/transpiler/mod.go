package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"fmt"
	"text/scanner"
)

func DoRoot(s *scanner.Scanner, output langoutput.T, src string) {
	// first scan
	s.Scan()

	for s.Peek() != scanner.EOF {
		tok := s.TokenText()
		s.Scan()
		switch tok {
		case "fonction":
			doFunction(s, output, src)
		case "procedure":
			doProcedure(s, output, src)
		default:
			panic(fmt.Sprintf("unexpected token '%s'", tok))
		}
	}
}
