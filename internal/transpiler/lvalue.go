package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler/translate"
	"strings"
)

func doDeclare(s scan.Scanner, output langoutput.T) {
	varNames := []string{s.Text()} // first variable name

	for s.Match(",") {
		varName := s.Text()
		varNames = append(varNames, varName)
	}

	s.Must(":")

	tabLength := doTypeMaybeSize(s, output)
	output.Writef(" %s", strings.Join(varNames, ", "))
	if tabLength != nil {
		output.Writef("(%s)", *tabLength)
	}

	// check if doing assignation at the same time
	if s.Match("<") {
		s.Must("-")
		value := s.Expr()
		s.Must(";")
		output.Writef(" = %s;", translate.Expr(value))
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

		value := s.Expr()
		s.Must(";")
		output.Writef("%s = %s;", lval, translate.Expr(value))
	} else if s.Match("(") { // function call
		doFunctionCall(s, output, lval)
	} else {
		s.InvalidToken("expected '<-' or '('")
	}
}

// reads the arguments of a function call **and no parenthesis**
// the first parenthesis must have been eaten
func getFunctionArgs(s scan.Scanner) []string {
	args := []string{}

	if s.Peek() == ")" {
		return args
	}

	for {
		arg := s.Expr()
		args = append(args, translate.Expr(arg))

		if s.Peek() == ")" {
			return args
		} else if s.Match(",") {
		} else {
			s.InvalidToken("expected ',' or ')'")
		}
	}
}

func doFunctionCall(s scan.Scanner, output langoutput.T, name string) {
	output.Writef("%s(", name)

	args := getFunctionArgs(s)
	output.Write(strings.Join(args, ", "))

	s.Must(")")
	s.Must(";")
	output.Write(");")
}
