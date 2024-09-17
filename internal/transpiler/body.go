package transpiler

import (
	"algo-iut-1/internal/transpiler/loops"
	"algo-iut-1/internal/transpiler/scanutils"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"io"
	"slices"
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
	if scanutils.Text(s) == "<" {
		if scanutils.Text(s) == "-" {
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

// insert spaces or tabs for code body lines
// returns "true" is we should process the line, "false" if EOF
func insertTabs(s *scanner.Scanner, output io.WriteCloser) bool {
	oldWhitespace := s.Whitespace
	// remove space and tab
	s.Whitespace ^= 1 << ' '
	s.Whitespace ^= 1 << '\t'

	for {
		if s.Scan() == scanner.EOF {
			return false
		}

		if slices.Contains([]string{" ", "\t"}, s.TokenText()) {
			output.Write([]byte(s.TokenText()))
		} else {
			break
		}
	}

	s.Whitespace = oldWhitespace
	return true
}

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s *scanner.Scanner, output io.WriteCloser) {
	for {
		if !insertTabs(s, output) {
			break
		}

		tok := s.TokenText()
		switch tok {
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
			doIdentifier(s, output, tok)
			// panic(fmt.Sprintf("Unknown token: '%s'", tok))
		}
		output.Write([]byte("\n"))

	}
}
