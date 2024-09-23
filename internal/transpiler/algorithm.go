package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
)

func doAlgorithme(s scan.Scanner, output langoutput.T, src string) {
	s.Advance() // ignore alg name
	s.Must("debut")

	output.Write("int main() {")
	doBody(s, output, src)
	// } is written by doBody()
}
