package transpiler

import (
	"algo-iut-1/internal/transpiler/scanutils"
	"algo-iut-1/internal/transpiler/translate"
	"io"
	"text/scanner"
)

func doFunctionHeader(s *scanner.Scanner, output io.WriteCloser) {
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

func doFunction(s *scanner.Scanner, output io.WriteCloser) {
	doFunctionHeader(s, output)
	doBody(s, output)
}
