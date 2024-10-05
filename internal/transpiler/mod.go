package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"fmt"
	"runtime/debug"
)

func Do(s scan.Scanner, output langoutput.T, src string) (success bool) {
	defer func() {
		if r := recover(); r != nil {
			showError(s, src, r)
			fmt.Println(string(debug.Stack()))
			success = false
		}
	}()

	output.Write("#include <iostream>\n")
	output.Write("#include <vector>\n")
	doRoot(s, output, src)

	return true
}

func doRoot(s scan.Scanner, output langoutput.T, src string) {
	for s.HasMore() {
		tok := s.Peek()
		switch tok {
		case "fonction":
			s.Advance()
			doFunction(s, output, src)
		case "procedure":
			s.Advance()
			doProcedure(s, output, src)
		case "algorithme":
			s.Advance()
			doAlgorithme(s, output, src)
		default:
			s.InvalidToken("")
		}
	}
}
