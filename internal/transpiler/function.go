package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler/translate"
)

func doFunctionHeader(s scan.Scanner, output langoutput.T) {
	// get function name
	functionName := s.Text()

	// get function args
	args := doFunctionOrProcedureArgs(s)

	// get return type
	s.Must("renvoie")
	retType := translate.Type(scanType(s))

	s.Must("debut")

	writeFunctionOrProcedureHeader(functionName, args, retType, output)
}

func doFunction(s scan.Scanner, output langoutput.T, src string) {
	doFunctionHeader(s, output)
	doBody(s, output, src)
}
