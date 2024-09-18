package transpiler

import (
	"algo-iut-1/internal/transpiler/loops"
	"algo-iut-1/internal/transpiler/scanutils"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"io"
	"text/scanner"
)

func doDeclare(s *scanner.Scanner, output io.WriteCloser) {
	varName := scanutils.Text(s)

	scanutils.Must(s, ":")

	varType, tabLength := translate.TypeMaybeSize(scanutils.UntilEOL(s))
	if tabLength == nil {
		output.Write([]byte(fmt.Sprintf("%v %v;", varType, varName)))
	} else {
		output.Write([]byte(fmt.Sprintf("%v %v(%v);", varType, varName, *tabLength)))
	}
}

// assume `x <-` is already scanned
func doAssignInner(s *scanner.Scanner, output io.WriteCloser, varName string) {
	value := scanutils.UntilEOL(s)

	output.Write([]byte(fmt.Sprintf("%s = %s;", varName, value)))
}

func doIdentifier(s *scanner.Scanner, output io.WriteCloser, id string) {
	if scanutils.Match(s, "<") {
		if scanutils.Match(s, "-") {
			doAssignInner(s, output, id)
		} else {
			PanicInvalidToken(s)
		}
	} else {
		PanicInvalidToken(s)
	}
}

func doReturn(s *scanner.Scanner, output io.WriteCloser) {
	value := scanutils.UntilEOL(s)
	output.Write([]byte(fmt.Sprintf("return %s;", value)))
}

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s *scanner.Scanner, output io.WriteCloser) {
	for {
		tok := s.TokenText()
		switch tok {
		case "fin":
			s.Scan()
			output.Write([]byte("}\n\n"))
			return
		case "declarer":
			s.Scan()
			doDeclare(s, output)
		case "renvoie":
			s.Scan()
			doReturn(s, output)
		case "pour":
			s.Scan()
			loops.DoPourLoop(s, output)
		default:
			s.Scan()
			doIdentifier(s, output, tok)
		}
		output.Write([]byte("\n"))

	}
}
