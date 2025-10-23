package loops

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler/translate"
)

func DoPourLoop(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	varName := s.Text()
	s.Must("variant_de")

	min := s.Expr()
	s.Must("a")
	max := s.Expr()

	s.Must(")")
	s.Must("faire")

	output.Write("for(int ")
	output.Write(varName)
	output.Write(" = ")
	output.Write(translate.Expr(min))
	output.Write("; ")
	output.Write(varName)
	output.Write(" <= ")
	output.Write(translate.Expr(max))
	output.Write("; ")
	output.Write(varName)
	output.Write("++) {")
}
