package transpiler

import (
	"algo-iut-1/internal/tabanalyser"
	"algo-iut-1/internal/transpiler/loops"
	"algo-iut-1/internal/transpiler/scanutils"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"io"
	"os"
	"strings"
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

// line that starts with an identifier. Identifier is already scanned as `id`
func doLValueStart(s *scanner.Scanner, output io.WriteCloser) {
	lval := scanutils.LValue(s)

	scanutils.Must(s, "<")
	scanutils.Must(s, "-")

	value := scanutils.UntilEOL(s)
	output.Write([]byte(fmt.Sprintf("%s = %s;", lval, value)))
}

func doReturn(s *scanner.Scanner, output io.WriteCloser) {
	value := scanutils.UntilEOL(s)
	output.Write([]byte(fmt.Sprintf("return %s;", value)))
}

func showError(s *scanner.Scanner, src string, errStr interface{}) {
	lines := strings.Split(src, "\n")
	line := lines[s.Pos().Line-1]

	fmt.Printf("Transpiler error: line %v\n", s.Pos().Line)
	fmt.Println(line)
	fmt.Println(strings.Repeat(" ", s.Pos().Column-1) + "^")
	fmt.Println(errStr)
}

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s *scanner.Scanner, output io.WriteCloser, src string) {
	tabsPrefix := tabanalyser.Do(src)
	for {
		if !doLine(s, output, tabsPrefix, src) {
			break
		}
	}
}

func doLine(s *scanner.Scanner, output io.WriteCloser, tabsPrefix []string, src string) bool {
	defer func() {
		if r := recover(); r != nil {
			showError(s, src, r)
			os.Exit(1)
		}
	}()

	// write tabs/space prefix
	prefix := tabsPrefix[s.Pos().Line-1]
	output.Write([]byte(prefix))

	tok := s.TokenText()
	switch tok {
	// conditions
	case "si":
		s.Scan()
		doCondition(s, output)
	// loops
	case "pour":
		s.Scan()
		loops.DoPourLoop(s, output)
	case "boucle":
		s.Scan()
		loops.DoInfiniteLoop(s, output)
	case "ffaire":
		s.Scan()
		output.Write([]byte("}"))
	// others
	case "declarer":
		s.Scan()
		doDeclare(s, output)
	case "renvoie":
		s.Scan()
		doReturn(s, output)
	case "fin":
		s.Scan()
		output.Write([]byte("}\n"))
		return false
	default:
		doLValueStart(s, output)
	}
	output.Write([]byte("\n"))

	return true
}
