package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/transpiler/scanutils"
	"algo-iut-1/internal/transpiler/translate"
	"text/scanner"
)

func doFunctionHeader(s *scanner.Scanner, output langoutput.T) {
	// get function name
	functionName := scanutils.Text(s)

	// get function args
	args := doFunctionOrProcedureArgs(s)

	// get return type
	scanutils.Must(s, "renvoie")
	retType := translate.Type(scanType(s))

	scanutils.Must(s, "debut")

	writeFunctionOrProcedureHeader(functionName, args, retType, output)
}

func doFunction(s *scanner.Scanner, output langoutput.T, src string) {
	doFunctionHeader(s, output)
	doBody(s, output, src)
}
