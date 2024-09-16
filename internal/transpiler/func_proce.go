package transpiler

import (
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"io"
	"text/scanner"
)

type typedVar struct {
	varType string
	varName string
	ref     bool
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

func doDeclare(s *scanner.Scanner, output io.WriteCloser) {
	s.Scan()
	varType := translate.Type(s.TokenText())

	s.Scan()
	varName := s.TokenText()

	output.Write([]byte(fmt.Sprintf("%v %v", varType, varName)))

	s.Scan()
	if s.TokenText() == "<" {
		mustScan(s, "-")

		s.Scan()
		varValue := s.TokenText()
		output.Write([]byte(fmt.Sprintf(" = %s", varValue)))

		s.Scan()
	}

	if s.TokenText() == ";" {
		output.Write([]byte(";"))
		return
	} else {
		panic(fmt.Sprintf("Invalid token: '%s'", s.TokenText()))
	}

}

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s *scanner.Scanner, output io.WriteCloser) {
	for s.Scan() != scanner.EOF {
		tok := s.TokenText()
		switch tok {
		case "fin":
			output.Write([]byte("}\n\n"))
			return
		case "declarer":
			doDeclare(s, output)
		default:
			panic(fmt.Sprintf("Unknown token: '%s'", tok))
		}

	}
}
