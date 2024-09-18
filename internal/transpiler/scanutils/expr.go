package scanutils

import (
	"slices"
	"text/scanner"
)

var operators = []string{"+", "-", "*", "/"}

const (
	ExprNextId       = iota
	ExprNextOperator = iota
)

// assume `(` is already eaten
func function(s *scanner.Scanner) []string {
	tokens := make([]string, 0)

	// handle empty params
	if s.TokenText() == ")" {
		s.Scan()
		return tokens
	}

	for {
		varName := s.TokenText()
		tokens = append(tokens, varName)
		s.Scan()

		if s.TokenText() == ")" {
			s.Scan()
			return tokens
		} else if s.TokenText() == "," {
			s.Scan()
		} else {
			panic("expected ','")
		}
	}
}

func Expr(s *scanner.Scanner) []string {
	tokens := make([]string, 0)

	id := s.TokenText()
	s.Scan()
	tokens = append(tokens, id)

	mode := ExprNextOperator

	for {
		isOperator := slices.Contains(operators, s.TokenText())

		if mode == ExprNextId { // if it expects an id
			if isOperator {
				panic("2 operators detected")
			} else {
				tokens = append(tokens, s.TokenText())
				s.Scan() // eat

				// check if its a function
				if s.TokenText() == "(" {
					s.Scan()
					tokens = append(tokens, function(s)...)
				}
			}
			mode = ExprNextOperator
		} else if mode == ExprNextOperator { // if it expects an operator

			if isOperator {
				tokens = append(tokens, s.TokenText())
				s.Scan() // eat
			} else {
				return tokens
			}
			mode = ExprNextId
		} else {
			panic("no")
		}
	}
}
