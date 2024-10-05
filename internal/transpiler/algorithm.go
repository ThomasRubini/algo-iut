package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
)

func doAlgorithme(s scan.Scanner, output langoutput.T, src string) {
	s.Advance() // ignore alg name
	s.Must("debut")

	output.Write("int main() {\n")
	doBody(s, output, src)
	// } is written by doBody()
}
