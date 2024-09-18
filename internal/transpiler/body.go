package transpiler

import (
	"algo-iut-1/internal/tabanalyser"
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

// line that starts with an identifier. Identifier is already scanned as `id`
func doIdentifierStart(s *scanner.Scanner, output io.WriteCloser, id string) {
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
func doBody(s *scanner.Scanner, output io.WriteCloser, src string) {
	tabsPrefix := tabanalyser.Do(src)

	for {
		// write tabs/space prefix
		prefix := tabsPrefix[s.Pos().Line-1]
		output.Write([]byte(prefix))

		tok := s.TokenText()
		s.Scan() // in any case it will be consumed
		switch tok {
		case "ffaire":
			output.Write([]byte("}\n\n"))
		case "fin":
			output.Write([]byte("}\n\n"))
			return
		case "declarer":
			doDeclare(s, output)
		case "renvoie":
			doReturn(s, output)
		case "pour":
			loops.DoPourLoop(s, output)
		default:
			doIdentifierStart(s, output, tok)
		}
		output.Write([]byte("\n"))

	}
}
