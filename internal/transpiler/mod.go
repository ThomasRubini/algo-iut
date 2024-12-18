package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"fmt"
	"runtime/debug"
)

func Do(s scan.Scanner, output langoutput.T, src string) (transpile_err *Error) {
	defer func() {
		if r := recover(); r != nil {
			// return error
			transpile_err = &Error{
				src:             src,
				errStr:          fmt.Sprintf("%v", r),
				s:               s,
				transpilerStack: string(debug.Stack()),
			}
		}
	}()

	output.Write("#include <iostream>\n")
	output.Write("#include <vector>\n")
	output.Write("\n")

	// We do not use `using namespace std;` to avoid collisions with a lot of function names people might write in the algo language (e.g. abs())
	output.Write("using std::cout, std::endl, std::cin, std::string, std::vector;\n")

	doRoot(s, output, src)

	return nil
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
