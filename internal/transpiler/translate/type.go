package translate

import (
	"algo-iut/internal/ref"
)

func Type(in string) *string {
	switch in {
	case "entier":
		return ref.String("int")
	case "entier_naturel":
		return ref.String("unsigned int")
	case "reel":
		return ref.String("double")
	case "char":
		return ref.String("char")
	case "string":
		return ref.String("std::string")
	case "booleen":
		return ref.String("bool")
	case "caractere":
		return ref.String("char")
	default:
		return nil
	}
}
