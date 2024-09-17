package transpiler

import (
	"algo-iut-1/internal/transpiler/scanutils"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"io"
	"slices"
	"text/scanner"
)

type typedVar struct {
	varType string
	varName string
	ref     bool
}

func doFunctionOrProcedureArgs(s *scanner.Scanner) []typedVar {
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

func writeFunctionOrProcedureHeader(functionName string, args []typedVar, retType string, output io.WriteCloser) {
	// write
	output.Write([]byte(fmt.Sprintf("%s %s(", retType, functionName)))
	for i, arg := range args {
		output.Write([]byte(fmt.Sprintf("%s %s", arg.varType, arg.varName)))
		if i < len(args)-1 {
			output.Write([]byte(", "))
		}
	}
	output.Write([]byte(") {\n"))
}
