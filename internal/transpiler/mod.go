package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
)

func DoRoot(s scan.Scanner, output langoutput.T, src string) {
	for s.HasMore() {
		tok := s.Peek()
		switch tok {
		case "fonction":
			s.Advance()
			doFunction(s, output, src)
		case "procedure":
			s.Advance()
			doProcedure(s, output, src)
		default:
			s.InvalidToken("")
		}
	}
}
