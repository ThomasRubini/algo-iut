package scan

import (
	"algo-iut/internal/ref"
	"algo-iut/internal/scan/scanexpr"
)

const (
	ExprNextId       = iota
	ExprNextOperator = iota
)

func tryGetOperator(s Scanner) *string {
	simpleOperators := "+-*/"
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

	if s.Match("vaut") {
		return ref.String("vaut")
	}

	// other condition operators
	if s.Match("ET_ALORS") {
		return ref.String("ET_ALORS")
	}

	if s.Match("OU_SINON") {
		return ref.String("OU_SINON")
	}
	return nil
}

// assume `(` is already eaten
func function(s Scanner, name string) scanexpr.CompFuncImpl {
	fun := scanexpr.CompFuncImpl{
		Name: name,
		Args: make([]scanexpr.Comp, 0),
	}

	// handle empty params
	if s.Match(")") {
		return fun
	}

	for {
		arg := s.Expr()
		fun.Args = append(fun.Args, arg)

		if s.Match(")") {
			return fun
		} else if s.Match(",") {
		} else {
			s.InvalidToken("expected ',' or ')'")
		}
	}
}

// id, array or function
// (id means variable or constant)
func idOrArrOrFun(s Scanner) scanexpr.Comp {
	id := s.Text()

	// check if its a function
	if s.Match("(") {
		return function(s, id)

	} else if s.Match("[") {
		index := s.Expr()
		s.Must("]")
		return scanexpr.CompArrImpl{
			Name:  id,
			Index: index,
		}
	} else {
		return scanexpr.CompIdImpl{
			Name: id,
		}
	}
}

func (s *impl) Expr() scanexpr.Comp {
	e := scanexpr.CompMergeImpl{
		Comps: make([]scanexpr.Comp, 0),
	}

	mode := ExprNextId

	bracketCount := 0
	for {
		// handle brackets ()
		if s.Match("(") {
			bracketCount++
			e.Comps = append(e.Comps, scanexpr.Op("("))
			continue
		}
		if bracketCount > 0 && s.Match(")") {
			bracketCount--
			e.Comps = append(e.Comps, scanexpr.Op(")"))
			continue
		}

		op := tryGetOperator(s)

		if mode == ExprNextId { // if it expects an id
			if op != nil {
				panic("2 operators detected")
			} else {
				e.Comps = append(e.Comps, idOrArrOrFun(s))
			}
			mode = ExprNextOperator
		} else if mode == ExprNextOperator { // if it expects an operator
			if op != nil {
				e.Comps = append(e.Comps, scanexpr.CompOpImpl{
					Op: *op,
				})
			} else {
				// return
				if len(e.Comps) == 1 {
					return e.Comps[0]
				} else {
					return e
				}
			}
			mode = ExprNextId
		} else {
			panic("no")
		}
	}
}
