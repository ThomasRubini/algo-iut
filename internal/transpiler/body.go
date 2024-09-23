package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/tabanalyser"
	"algo-iut-1/internal/transpiler/loops"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"os"
	"strings"
)

func doDeclare(s scan.Scanner, output langoutput.T) {
	varName := s.Text()

	s.Must(":")

	varType, tabLength := translate.TypeMaybeSize(s.UntilEOL())
	if tabLength == nil {
		output.Write(fmt.Sprintf("%v %v;", varType, varName))
	} else {
		output.Write(fmt.Sprintf("%v %v(%v);", varType, varName, *tabLength))
	}
}

// line that starts with an identifier. Identifier is already scanned as `id`
func doLValueStart(s scan.Scanner, output langoutput.T) {
	lval := s.LValue()

	s.Must("<")
	s.Must("-")

	value := s.UntilEOL()
	output.Writef("%s = %s;", lval, value)
}

func doReturn(s scan.Scanner, output langoutput.T) {
	value := s.UntilEOL()
	output.Writef("return %s;", value)
}

func showError(s scan.Scanner, src string, errStr interface{}) {
	lines := strings.Split(src, "\n")
	line := lines[s.Pos().Line-1]

	fmt.Printf("Transpiler error: line %v\n", s.Pos().Line)
	fmt.Println(line)
	fmt.Println(strings.Repeat(" ", s.Pos().Column-1) + "^")
	fmt.Println(errStr)
}

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s scan.Scanner, output langoutput.T, src string) {
	tabsPrefix := tabanalyser.Do(src)
	for {
		if !doLine(s, output, tabsPrefix, src) {
			break
		}
	}
}

func doLine(s scan.Scanner, output langoutput.T, tabsPrefix []string, src string) bool {
	defer func() {
		if r := recover(); r != nil {
			showError(s, src, r)
			os.Exit(1)
		}
	}()

	// write tabs/space prefix
	prefix := tabsPrefix[s.Pos().Line-1]
	output.Write(prefix)

	tok := s.Peek()
	switch tok {
	// conditions
	case "si":
		s.Advance()
		doCondition(s, output)
	// loops
	case "pour":
		s.Advance()
		loops.DoPourLoop(s, output)
	case "boucle":
		s.Advance()
		loops.DoInfiniteLoop(s, output)
	case "ffaire":
		s.Advance()
		output.Write("}")
	// others
	case "declarer":
		s.Advance()
		doDeclare(s, output)
	case "renvoie":
		s.Advance()
		doReturn(s, output)
	case "fin":
		s.Advance()
		output.Write("}\n")
		return false
	default:
		doLValueStart(s, output)
	}
	output.Write("\n")

	return true
}
