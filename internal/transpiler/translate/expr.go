package translate

import (
	"algo-iut-1/internal/scan/scanexpr"
	"fmt"
	"strings"
)

func exprFunction(e scanexpr.CompFuncImpl) string {
	if e.Name == "taille" {
		if len(e.Args) != 1 {
			panic("taille() must have exactly one argument")
		}
		return fmt.Sprintf("%s.size()", Expr(e.Args[0]))
	}

	args := make([]string, len(e.Args))
	for i, arg := range e.Args {
		args[i] = Expr(arg)
	}
	return e.Name + "(" + strings.Join(args, ", ") + ")"
}

func Expr(e scanexpr.Comp) string {
	switch e.Type() {
	case scanexpr.CompVar:
		return e.(scanexpr.CompVarImpl).Name
	case scanexpr.CompFunc:
		return exprFunction(e.(scanexpr.CompFuncImpl))
	case scanexpr.CompArr:
		arr := e.(scanexpr.CompArrImpl)
		return arr.Name + "[" + Expr(arr.Index) + "]"
	case scanexpr.CompOp:
		return Operator(e.(scanexpr.CompOpImpl).Op)
	case scanexpr.CompMerge:
		merge := e.(scanexpr.CompMergeImpl)
		comps := make([]string, len(merge.Comps))
		for i, comp := range merge.Comps {
			comps[i] = Expr(comp)
		}
		return strings.Join(comps, " ")
	}

	panic("unknown expression type")
}
