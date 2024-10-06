package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/ref"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler/translate"
)

// type as a function argument (without size possible)
func doTypeNoSize(s scan.Scanner, output langoutput.T) {
	size := doTypeMaybeSize(s, output)
	if size != nil {
		panic("Found size in type when disallowed")
	}
}

func doTypeMaybeSize(s scan.Scanner, output langoutput.T) (size *string) {
	// TODO this is a hack. Doesn't properly support nested tableaux
	if s.Match("tableau_de") {
		tok := s.Peek()
		maybeType := translate.Type(tok)
		if maybeType == nil { // if not a type, then it must be a size
			size := s.Expr()
			output.Write("std::vector<")
			_ = doTypeMaybeSize(s, output)
			output.Write(">")
			return ref.String(translate.Expr(size))
		} else {
			output.Writef("std::vector<%s>", *maybeType)
			return nil
		}
	} else if s.Match("constante") {
		output.Write("const ")
		return doTypeMaybeSize(s, output)
	} else {
		typ := translate.Type(s.Text())
		if typ == nil {
			panic("Unknown type: " + s.Text())
		}
		output.Write(*typ)
		return nil
	}
}
