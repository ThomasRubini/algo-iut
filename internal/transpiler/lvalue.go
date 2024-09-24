package transpiler

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler/translate"
	"fmt"
)

func doDeclare(s scan.Scanner, output langoutput.T) {
	varName := s.Text()

	s.Must(":")

	varType, tabLength := translate.TypeMaybeSize(s)
	if tabLength == nil {
		output.Write(fmt.Sprintf("%v %v", varType, varName))
	} else {
		output.Write(fmt.Sprintf("%v %v(%v)", varType, varName, *tabLength))
	}

	// check if doing assignation at the same time
	if s.Match("<") {
		s.Must("-")
		value := s.UntilEOL() // eats the ';'
		output.Writef(" = %s;", value)
	} else {
		s.Must(";")
	}
	output.Write(";")
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
		} else if s.Match(",") {
		} else {
			s.InvalidToken("expected ',' or ')'")
		}
	}
}
