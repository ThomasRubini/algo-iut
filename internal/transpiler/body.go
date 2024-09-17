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
	s.Scan()
	varType := translate.Type(s.TokenText())

	s.Scan()
	varName := s.TokenText()

	output.Write([]byte(fmt.Sprintf("%v %v", varType, varName)))

	s.Scan()
	if s.TokenText() == "<" {
		scanutils.Must(s, "-")

		s.Scan()
		varValue := s.TokenText()
		output.Write([]byte(fmt.Sprintf(" = %s", varValue)))

		s.Scan()
	}

	if s.TokenText() == ";" {
		output.Write([]byte(";"))
		return
	} else {
		panic(fmt.Sprintf("Invalid token: '%s'", s.TokenText()))
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

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s *scanner.Scanner, output io.WriteCloser) {
	for s.Scan() != scanner.EOF {
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
