package loops

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
)

func DoInfiniteLoop(s scan.Scanner, output langoutput.T) {
	output.Write("while(true) {")
}

func DoBreak(s scan.Scanner, output langoutput.T) {
	s.Must(";")
	output.Write("break;")
}

func DoWhile(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	cond := s.Expr()
	s.Must(")")
	output.Writef("while(%v) {", cond)
}
