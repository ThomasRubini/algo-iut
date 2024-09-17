package translate

import (
	"algo-iut-1/internal/ref"
	"fmt"
	"strings"
)

// translates a type into a C++ type, and returns its size.
func TypeMaybeSize(in string) (string, *string) {
	// TODO this is a hack. Doesn't support nested tableaux
	parts := strings.Split(in, " ")
	if parts[0] == "tableau_de" {
		if len(parts) == 2 {
			return fmt.Sprintf("std::vector<%s>", Type(parts[1])), ref.String("0")
		} else {
			size := strings.Join(parts[1:len(parts)-1], " ")
			typeName := parts[len(parts)-1]
			return fmt.Sprintf("std::vector<%s>", Type(typeName)), &size
		}
	} else {
		return Type(in), nil
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
