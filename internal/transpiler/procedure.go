package transpiler

import (
	"algo-iut-1/internal/transpiler/scanutils"
	"io"
	"text/scanner"
)

func doProcedureHeader(s *scanner.Scanner, output io.WriteCloser) {
	// get function name
	functionName := scanutils.Text(s)

	// get function args
	args := doFunctionOrProcedureArgs(s)

	scanutils.Must(s, "debut")

	writeFunctionOrProcedureHeader(functionName, args, "void", output)
}

func doProcedure(s *scanner.Scanner, output io.WriteCloser) {
	doProcedureHeader(s, output)
	doBody(s, output)
}
