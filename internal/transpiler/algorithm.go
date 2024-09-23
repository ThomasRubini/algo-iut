package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
)

func doAlgorithme(s scan.Scanner, output langoutput.T, src string) {
	s.Advance() // ignore alg name
	s.Must("debut")

	output.Write("void main() {")
	doBody(s, output, src)
	output.Write("}")
}
