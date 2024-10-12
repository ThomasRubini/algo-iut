package translate

import (
	"algo-iut/internal/scan/scanexpr"
	"fmt"
	"strings"
)

func exprFunction(e scanexpr.CompFuncImpl) string {
	// try special functions
	switch e.Name {
	case "succ":
		if len(e.Args) != 1 {
			panic("succ() must have exactly one argument")
		}
		return fmt.Sprintf("(char) (++%s)", Expr(e.Args[0]))
	case "prec":
		if len(e.Args) != 1 {
			panic("prec() must have exactly one argument")
		}
		return fmt.Sprintf("(char) (--%s)", Expr(e.Args[0]))
	case "modulo":
		if len(e.Args) != 2 {
			panic("modulo() must have exactly 2 arguments")
		}
		return fmt.Sprintf("%s %% %s", Expr(e.Args[0]), Expr(e.Args[1]))
	case "taille":
		if len(e.Args) != 1 {
			panic("taille() must have exactly one argument")
		}
		return fmt.Sprintf("%s.size()", Expr(e.Args[0]))
	case "rand":
		if len(e.Args) != 2 {
			panic("rand() must have exactly 2 argument")
		}
		return fmt.Sprintf("rand() %% %s + %s", Expr(e.Args[1]), Expr(e.Args[0]))
	case "rang":
		if len(e.Args) != 1 {
			panic("rang() must have exactly 1 argument")
		}
		return fmt.Sprintf("(char) (%s)", Expr(e.Args[0]))
	}

	args := make([]string, len(e.Args))
	for i, arg := range e.Args {
		args[i] = Expr(arg)
	}
	return e.Name + "(" + strings.Join(args, ", ") + ")"
}

func Expr(e scanexpr.Comp) string {
	switch e.Type() {
	case scanexpr.CompId:
		return e.(scanexpr.CompIdImpl).Name
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
