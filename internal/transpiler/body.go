package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/tabanalyser"
	"algo-iut/internal/transpiler/loops"
	"algo-iut/internal/transpiler/translate"
	"strings"
)

func doReturn(s scan.Scanner, output langoutput.T) {
	value := translate.Expr(s.Expr())
	s.Must(";")
	output.Writef("return %s;", value)
}

func doAfficher(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	args := getFunctionArgs(s)
	s.Must(")")
	s.Must(";")
	output.Writef("std::cout << %s;", strings.Join(args, " << "))
}

func doSaisir(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	variable := s.LValue()
	s.Must(")")
	s.Must(";")
	output.Writef("std::cin >> %s;", variable)
}

func doAllonger(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	vec := translate.Expr(s.Expr())
	s.Must(",")
	amount := translate.Expr(s.Expr())
	s.Must(")")
	s.Must(";")

	output.Writef("%v.resize(%v.size() + %v);", vec, vec, amount)
}
func doRedimensionner(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	vec := translate.Expr(s.Expr())
	s.Must(",")
	amount := translate.Expr(s.Expr())
	s.Must(")")
	s.Must(";")

	output.Writef("%v.resize(%v);", vec, amount)
}

func doPermuter(s scan.Scanner, output langoutput.T) {
	s.Must("(")
	valueA := translate.Expr(s.Expr())
	s.Must(",")
	valueB := translate.Expr(s.Expr())
	s.Must(")")
	s.Must(";")

	output.Write("{")
	output.Writef("auto tmp = %v; ", valueA)
	output.Writef("%v = %v; ", valueA, valueB)
	output.Writef("%v = tmp; ", valueB)
	output.Write("}")
}

// scan a function/procedure body. Returns when encountering "fin"
func doBody(s scan.Scanner, output langoutput.T, src string) {
	tabsPrefix := tabanalyser.Do(src)
	for {
		if !doLine(s, output, tabsPrefix) {
			break
		}
	}
}

func doLigneSuivante(s scan.Scanner, output langoutput.T) {
	s.Must(";")
	output.Write("std::cout << std::endl;")
}

func doLine(s scan.Scanner, output langoutput.T, tabsPrefix []string) bool {
	// write tabs/space prefix
	prefix := tabsPrefix[s.Pos().Line-1]
	output.Write(prefix)

	tok := s.Peek()
	switch tok {
	// conditions
	case "si":
		s.Advance()
		doCondition(s, output)
	case "sinon":
		s.Advance()
		output.Write("} else {")
	case "sinon_si":
		s.Advance()
		doElseIf(s, output)
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
	case "continue":
		s.Advance()
		loops.DoContinue(s, output)
	case "sortie":
		s.Advance()
		loops.DoBreak(s, output)
	case "ffaire":
		s.Advance()
		output.Write("}")
	case "fboucle":
		s.Advance()
		output.Write("}")
	// special functions
	case "afficher":
		s.Advance()
		doAfficher(s, output)
	case "saisir":
		s.Advance()
		doSaisir(s, output)
	case "allonger":
		s.Advance()
		doAllonger(s, output)
	case "redimensionner":
		s.Advance()
		doRedimensionner(s, output)
	case "permuter":
		s.Advance()
		doPermuter(s, output)
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
