package scanutils

import (
	"algo-iut-1/internal/ref"
	"fmt"
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
		varName := Text(s)
		tokens = append(tokens, varName)

		if Match(s, ")") {
			fmt.Printf("C (next is %v)\n", s.TokenText())
			return tokens
		} else if Match(s, ",") {
		} else {
			panic("expected ','")
		}
	}
}

func TextOrFunction(s *scanner.Scanner) []string {
	fmt.Printf("D (next is %v)\n", s.TokenText())
	id := Text(s)

	// check if its a function
	if Match(s, "(") {
		l := make([]string, 0)
		l = append(l, id)
		l = append(l, function(s)...)
		return l
	} else {
		return []string{id}
	}
}

func Expr(s *scanner.Scanner) []string {
	fmt.Printf("Call to Expr() (next is %v)\n", s.TokenText())
	tokens := make([]string, 0)

	tokens = append(tokens, TextOrFunction(s)...)

	mode := ExprNextOperator

	for {
		op := tryGetOperator(s)

		if mode == ExprNextId { // if it expects an id
			fmt.Printf("Want ID (next is %v)\n", s.TokenText())
			if op != nil {
				panic("2 operators detected")
			} else {
				tokens = append(tokens, TextOrFunction(s)...)
				fmt.Printf("A (next is %v)\n", s.TokenText())
			}
			mode = ExprNextOperator
		} else if mode == ExprNextOperator { // if it expects an operator
			fmt.Printf("Want operator (next is %v)\n", s.TokenText())
			if op != nil {
				tokens = append(tokens, *op)
			} else {
				fmt.Printf("RETURN (next is %v)\n", s.TokenText())
				return tokens
			}
			mode = ExprNextId
		} else {
			panic("no")
		}
	}
}
