package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/tabanalyser"
	"algo-iut-1/internal/transpiler/loops"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
	"os"
	"runtime/debug"
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

	if s.Match("<") { // assignation
		s.Must("-")

		value := s.UntilEOL()
		output.Writef("%s = %s;", lval, value)
	} else if s.Match("(") { // function call
		doFunctionCall(s, output, lval)
	}
}

func doFunctionCall(s scan.Scanner, output langoutput.T, name string) {
	output.Writef("%s(", name)

	if s.Match(")") {
		output.Write(");")
		return
	}

	for {
		arg := s.Expr()
		output.Write(translate.Expr(arg))

		if s.Match(")") {
			s.Must(";")
			output.Write(");")
			break
		}

		s.Must(",")
	}
}

func doReturn(s scan.Scanner, output langoutput.T) {
	value := s.UntilEOL()
	output.Writef("return %s;", value)
}

func doAfficher(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	value := translate.Expr(s.Expr())
	s.Must(")")
	output.Writef("std::cout << %s << std::endl;", value)
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
			fmt.Println(string(debug.Stack()))
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
	case "fsi":
		s.Advance()
		output.Write("}")
	// loops
	case "pour":
		s.Advance()
		loops.DoPourLoop(s, output)
	case "boucle":
		s.Advance()
		loops.DoInfiniteLoop(s, output)
	case "tant_que":
		s.Advance()
		loops.DoWhile(s, output)
	case "sortie":
		s.Advance()
		loops.DoBreak(s, output)
	case "ffaire":
		s.Advance()
		output.Write("}")
	case "fboucle":
		s.Advance()
		output.Write("}")
	// afficher special case
	case "afficher":
		s.Advance()
		doAfficher(s, output)
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
