package transpiler

import (
	"algo-iut-1/internal/transpiler/scanutils"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"io"
	"slices"
	"text/scanner"
)

func doProcedureArgs(s *scanner.Scanner) []typedVar {
	scanutils.Must(s, "(")

	args := make([]typedVar, 0)

	// handle empty args
	s.Scan()
	if s.TokenText() == ")" {
		return args
	}

	for {
		// get var name
		varName := s.TokenText()

		scanutils.Must(s, ":")

		// check arg type
		var needRef bool
		s.Scan()
		argType := s.TokenText()
		if slices.Contains([]string{"in", "out"}, argType) {
			// idk if there is a real difference between them in generated C++
			needRef = argType == "out"
		} else {
			panic(fmt.Sprintf("Invalid arg type: '%s'", argType))
		}

		// get var type
		varType := translate.Type(scanType(s))

		// append
		args = append(args, typedVar{varType, varName, needRef})

		// check for end/next arg
		s.Scan()
		if s.TokenText() == ")" {
			return args
		} else if s.TokenText() == "," {
			s.Scan() // prepare for next scan at the beginning of the loop
			continue
		} else {
			panic(fmt.Sprintf("expected , or ), got %s", s.TokenText()))
		}
	}

}

func doProcedureHeader(s *scanner.Scanner, output io.WriteCloser) {
	// get function name
	s.Scan()
	functionName := s.TokenText()

	// get function args
	args := doProcedureArgs(s)

	scanutils.Must(s, "debut")

	writeFunctionOrProcedureHeader(functionName, args, "void", output)
}

func doProcedure(s *scanner.Scanner, output io.WriteCloser) {
	doProcedureHeader(s, output)
	doBody(s, output)
}
