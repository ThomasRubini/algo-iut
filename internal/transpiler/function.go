package transpiler

import (
	"algo-iut-1/internal/transpiler/scanutils"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"io"
	"text/scanner"
)

// transpile a function arguments
func doFunctionArgs(s *scanner.Scanner) []typedVar {
	scanutils.Must(s, "(")

	args := make([]typedVar, 0)

	// handle empty args
	s.Scan()
	if s.TokenText() == ")" {
		return args
	}

	for {
		// get var type
		varType := translate.Type(s.TokenText())

		// get var name
		s.Scan()
		varName := s.TokenText()

		// append
		args = append(args, typedVar{varType, varName, false})

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

func doFunctionHeader(s *scanner.Scanner, output io.WriteCloser) {
	// get function name
	s.Scan()
	functionName := s.TokenText()

	// get function args
	args := doFunctionArgs(s)

	// get return type
	scanutils.Must(s, "renvoie")
	retType := translate.Type(scanType(s))

	scanutils.Must(s, "debut")

	writeFunctionOrProcedureHeader(functionName, args, retType, output)
}

func doFunction(s *scanner.Scanner, output io.WriteCloser) {
	doFunctionHeader(s, output)
	doBody(s, output)
}
