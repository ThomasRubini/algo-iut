package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler/translate"
	"fmt"
)

func doCondition(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	condition := translate.Expr(s.Expr())
	s.Must(")")

	output.Write(fmt.Sprintf("if (%v) {", condition))
}
