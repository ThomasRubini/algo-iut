package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"slices"
)

type typedVar struct {
	varType string
	varName string
	ref     bool
}

func doFunctionOrProcedureArgs(s scan.Scanner) []typedVar {
	s.Must("(")

	args := make([]typedVar, 0)

	// handle empty args
	if s.Match(")") {
		return args
	}

	for {
		// get var name
		varName := s.Text()

		s.Must(":")

		// check arg type
		var needRef bool
		argType := s.Peek()
		if slices.Contains([]string{"in", "out"}, argType) {
			s.Advance()
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
		if s.Peek() == ")" {
			s.Advance()
			return args
		} else if s.Peek() == "," {
			s.Advance()
			continue
		} else {
			panic(fmt.Sprintf("expected , or ), got %s", s.Peek()))
		}
	}

}

func writeFunctionOrProcedureHeader(functionName string, args []typedVar, retType string, output langoutput.T) {
	// write
	output.Writef("%s %s(", retType, functionName)
	for i, arg := range args {
		output.Write(fmt.Sprintf("%s %s", arg.varType, arg.varName))
		if i < len(args)-1 {
			output.Write(", ")
		}
	}
	output.Write(") {\n")
}
