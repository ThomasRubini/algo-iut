package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"fmt"
	"strings"
)

func doCondition(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	condition := s.Expr()
	s.Must(")")

	output.Write(fmt.Sprintf("if (%v) {", strings.Join(condition, " ")))
}
