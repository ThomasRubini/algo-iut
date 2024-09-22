package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/transpiler/scanutils"
	"text/scanner"
)

func doProcedureHeader(s *scanner.Scanner, output langoutput.T) {
	// get function name
	functionName := scanutils.Text(s)

	// get function args
	args := doFunctionOrProcedureArgs(s)

	scanutils.Must(s, "debut")

	writeFunctionOrProcedureHeader(functionName, args, "void", output)
}

func doProcedure(s *scanner.Scanner, output langoutput.T, src string) {
	doProcedureHeader(s, output)
	doBody(s, output, src)
}
