package translate

import (
	"fmt"
	"strings"
)

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
