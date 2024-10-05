package translate

import (
	"algo-iut/internal/ref"
	"algo-iut/internal/scan"
	"algo-iut/internal/utils"
	"fmt"
	"strings"
)

// translates a type into a C++ type, and returns its size.
func TypeMaybeSize(s scan.Scanner) (string, *string) {
	// TODO this is a hack. Doesn't support nested tableaux
	if s.Match("tableau_de") {
		tok := s.Peek()
		typ1, err := utils.Catch(func() string {
			return Type(tok)
		})
		if err != nil { // if not a type, then it must be a size
			size := s.Expr()
			typ2, _ := TypeMaybeSize(s)
			return fmt.Sprintf("std::vector<%s>", typ2), ref.String(Expr(size))
		} else {
			return fmt.Sprintf("std::vector<%s>", typ1), ref.String("0")
		}
	} else {
		return Type(s.Text()), nil
	}
}

func Type(in string) string {
	after, found := strings.CutPrefix(in, "tableau_de ")
	if found {
		return fmt.Sprintf("std::vector<%s>", Type(after))
	}

	switch in {
	case "entier":
		return "int"
	case "entier_naturel":
		return "unsigned int"
	case "reel":
		return "double"
	case "char":
		return "char"
	case "string":
		return "string"
	case "booleen":
		return "bool"
	default:
		panic(fmt.Sprintf("unknown type %s", in))
	}
}
