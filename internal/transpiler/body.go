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

func doReturn(s scan.Scanner, output langoutput.T) {
	value := s.UntilEOL()
	output.Writef("return %s;", value)
}

func doAfficher(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	value := translate.Expr(s.Expr())
	s.Must(")")
	s.Must(";")
	output.Writef("std::cout << %s << std::endl;", value)
}

func showError(s scan.Scanner, src string, errStr interface{}) {
	lines := strings.Split(src, "\n")
	line := lines[s.Pos().Line-1]

	fmt.Printf("Transpiler error: line %v\n", s.Pos().Line)
	fmt.Println(line)
	fmt.Print(strings.Repeat(" ", s.Pos().Column+1) + "^")
	fmt.Println(strings.Repeat("-", len(s.Peek())-1))
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

func doLigneSuivante(s scan.Scanner, output langoutput.T) {
	s.Must(";")
	output.Write("std::cout << std::endl;")
}

func doLine(s scan.Scanner, output langoutput.T, tabsPrefix []string, src string) bool {
	defer func() {
		if r := recover(); r != nil {
			showError(s, src, r)
			fmt.Println(string(debug.Stack()))
			os.Exit(2)
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
	case "tant_que":
		s.Advance()
		loops.DoWhile(s, output)
	case "repeter":
		s.Advance()
		loops.DoRepeat(s, output)
	case "jusqua":
		s.Advance()
		loops.DoUntil(s, output)
	case "repeat":
		s.Advance()
		loops.DoRepeatUntil(s, output)
	case "boucle":
		s.Advance()
		loops.DoInfiniteLoop(s, output)
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
	case "ligne_suivante":
		s.Advance()
		doLigneSuivante(s, output)
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
