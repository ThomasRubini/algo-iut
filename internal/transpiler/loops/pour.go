package loops

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"strings"
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

	output.Writef("for(int %v=%v;i<%v;i++) {", varName, strings.Join(min, " "), strings.Join(max, " "))
}
