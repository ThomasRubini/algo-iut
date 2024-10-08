package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
)

func doFunctionHeader(s scan.Scanner, output langoutput.T) {
	// get function name
	functionName := s.Text()

	// get function args
	args := getFunctionOrProcedureHeaderArgs(s)

	// get return type
	s.Must("renvoie")
	retTypeOutput := langoutput.NewString()
	doTypeNoSize(s, retTypeOutput)
	s.Must("debut")

	writeFunctionOrProcedureHeader(functionName, args, retTypeOutput.String(), output)
}

func doFunction(s scan.Scanner, output langoutput.T, src string) {
	doFunctionHeader(s, output)
	doBody(s, output, src)
}
