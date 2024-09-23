package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
)

func doCondition(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	condition := translate.Expr(s.Expr())
	s.Must(")")

	output.Write(fmt.Sprintf("if (%v) {", condition))
}
