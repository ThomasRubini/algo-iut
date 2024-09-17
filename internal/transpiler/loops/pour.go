package loops

import (
	"algo-iut-1/internal/transpiler/scanutils"
	"fmt"
	"io"
	"strings"
	"text/scanner"
)

func DoPourLoop(s *scanner.Scanner, output io.WriteCloser) {
	scanutils.Must(s, "(")
	varName := scanutils.Text(s)
	scanutils.Must(s, "variant_de")

	min := scanutils.Expr(s)

	if s.TokenText() != "a" {
		panic("no")
	}
	// scanutils.Must(s, "a")
	max := scanutils.Expr(s)

	// scanutils.Must(s, ")");
	if s.TokenText() != ")" {
		panic("no")
	}

	output.Write([]byte(fmt.Sprintf("for(int %v=%v;i<%v;i++) {", varName, strings.Join(min, " "), strings.Join(max, " "))))
}
