package transpiler

import (
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

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s *scanner.Scanner, output io.WriteCloser) {
	for s.Scan() != scanner.EOF {
		tok := s.TokenText()
		if tok == "fin" {
			output.Write([]byte("}\n\n"))
			return
		}
	}
}
