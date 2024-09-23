package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"fmt"
)

func DoRoot(s scan.Scanner, output langoutput.T, src string) {
	for s.HasMore() {
		tok := s.Peek()
		s.Advance()
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
