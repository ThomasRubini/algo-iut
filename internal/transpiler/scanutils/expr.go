package scanutils

import (
	"algo-iut-1/internal/ref"
	"algo-iut-1/internal/utils"
	"text/scanner"
)

const (
	ExprNextId       = iota
	ExprNextOperator = iota
)

func tryGetOperator(s *scanner.Scanner) *string {
	var simpleOperators = "+-*/"
	for _, c := range simpleOperators {
		if Match(s, string(c)) {
			return ref.String(string(c))
		}
	}

	if Match(s, "=") {
		Must(s, "=")
		return ref.String("==")
	}
	if Match(s, "!") {
		Must(s, "=")
		return ref.String("!=")
	}
	if Match(s, ">") {
		Match(s, "=")
		return ref.String(">=")
	}
	if Match(s, "<") {
		Match(s, "=")
		return ref.String("<=")
	}
	return nil
}

// assume `(` is already eaten
func function(s *scanner.Scanner) []string {
	tokens := make([]string, 0)

	// handle empty params
	if Match(s, ")") {
		return tokens
	}

	for {
		varName := Expr(s)
		tokens = append(tokens, varName...)

		if Match(s, ")") {
			return tokens
		} else if Match(s, ",") {
			tokens = append(tokens, ",")
		} else {
			utils.PanicInvalidToken(s, "expected ',' or ')'")
		}
	}
}

func TextOrFunction(s *scanner.Scanner) []string {
	id := Text(s)

	// check if its a function
	if Match(s, "(") {
		l := make([]string, 0)
		l = append(l, id)
		l = append(l, "(")
		l = append(l, function(s)...)
		l = append(l, ")")
		return l
	} else {
		return []string{id}
	}
}

func Expr(s *scanner.Scanner) []string {
	tokens := make([]string, 0)

	tokens = append(tokens, TextOrFunction(s)...)

	mode := ExprNextOperator

	for {
		op := tryGetOperator(s)

		if mode == ExprNextId { // if it expects an id
			if op != nil {
				panic("2 operators detected")
			} else {
				tokens = append(tokens, TextOrFunction(s)...)
			}
			mode = ExprNextOperator
		} else if mode == ExprNextOperator { // if it expects an operator
			if op != nil {
				tokens = append(tokens, *op)
			} else {
				return tokens
			}
			mode = ExprNextId
		} else {
			panic("no")
		}
	}
}
