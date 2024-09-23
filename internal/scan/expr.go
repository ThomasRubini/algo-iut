package scan

import (
	"algo-iut-1/internal/ref"
)

const (
	ExprNextId       = iota
	ExprNextOperator = iota
)

func tryGetOperator(s Scanner) *string {
	var simpleOperators = "+-*/"
	for _, c := range simpleOperators {
		if s.Match(string(c)) {
			return ref.String(string(c))
		}
	}

	if s.Match("=") {
		s.Must("=")
		return ref.String("==")
	}
	if s.Match("!") {
		s.Must("=")
		return ref.String("!=")
	}
	if s.Match(">") {
		s.Match("=")
		return ref.String(">=")
	}

	if s.Match("<") {
		s.Match("=")
		return ref.String("<=")
	}

	// fully-text operators
	if s.Match("ne_vaut_pas") {
		return ref.String("ne_vaut_pas")
	}
	return nil
}

// assume `(` is already eaten
func function(s Scanner) []string {
	tokens := make([]string, 0)

	// handle empty params
	if s.Match(")") {
		return tokens
	}

	for {
		varName := s.Expr()
		tokens = append(tokens, varName...)

		if s.Match(")") {
			return tokens
		} else if s.Match(",") {
			tokens = append(tokens, ",")
		} else {
			s.InvalidToken("expected ',' or ')'")
		}
	}
}

// variable, array or function
func varOrArrOrFun(s Scanner) []string {
	id := s.Text()

	// check if its a function
	if s.Match("(") {
		l := make([]string, 0)
		l = append(l, id)
		l = append(l, "(")
		l = append(l, function(s)...)
		l = append(l, ")")
		return l

	} else if s.Match("[") {
		l := make([]string, 0)
		l = append(l, id)
		l = append(l, "[")
		l = append(l, s.Expr()...)
		l = append(l, "]")
		s.Must("]")
		return l
	} else {
		return []string{id}
	}
}

func (s *impl) Expr() []string {
	tokens := make([]string, 0)

	tokens = append(tokens, varOrArrOrFun(s)...)

	mode := ExprNextOperator

	for {
		op := tryGetOperator(s)

		if mode == ExprNextId { // if it expects an id
			if op != nil {
				panic("2 operators detected")
			} else {
				tokens = append(tokens, varOrArrOrFun(s)...)
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
