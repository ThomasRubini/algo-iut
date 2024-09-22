package transpiler

import (
	"algo-iut-1/internal/transpiler/scanutils"
	"fmt"
	"io"
	"text/scanner"
)

func doCondition(s *scanner.Scanner, output io.WriteCloser) {
	scanutils.Must(s, "(")
	condition := scanutils.Expr(s)
	scanutils.Must(s, ")")

	output.Write([]byte(fmt.Sprintf("if (%v) {", condition)))
}
