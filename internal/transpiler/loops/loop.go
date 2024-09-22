package loops

import (
	"algo-iut-1/internal/langoutput"
	"text/scanner"
)

func DoInfiniteLoop(s *scanner.Scanner, output langoutput.T) {
	output.Write("while(true) {")
}
