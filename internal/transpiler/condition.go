package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/transpiler/scanutils"
	"fmt"
	"strings"
	"text/scanner"
)

func doCondition(s *scanner.Scanner, output langoutput.T) {
	scanutils.Must(s, "(")
	condition := scanutils.Expr(s)
	scanutils.Must(s, ")")

	output.Write(fmt.Sprintf("if (%v) {", strings.Join(condition, " ")))
}
