package loops

import (
	"algo-iut-1/internal/transpiler/scanutils"
	"fmt"
	"io"
	"text/scanner"
)

func DoPourLoop(s *scanner.Scanner, output io.WriteCloser) {
	scanutils.Must(s, "(")
	varName := scanutils.Text(s)
	scanutils.Must(s, "variant_de")

	min := scanutils.Number(s)
	scanutils.Must(s, "a")
	max := scanutils.Number(s)

	scanutils.Must(s, ")");

	output.Write([]byte(fmt.Sprintf("for(int %v=%v;i<%v;i++) {", varName, min, max)))
}
