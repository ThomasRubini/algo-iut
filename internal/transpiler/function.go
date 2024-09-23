package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler/translate"
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
