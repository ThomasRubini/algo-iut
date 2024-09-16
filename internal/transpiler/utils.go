package transpiler

import (
	"fmt"
	"text/scanner"
)

func mustScan(s *scanner.Scanner, str string) {
	s.Scan()
	if s.TokenText() != str {
		panic(fmt.Sprintf("expected %s, got %s", str, s.TokenText()))
	}
}

func scanType(s *scanner.Scanner) string {
	tab_type := ""

	for {
		s.Scan()
		if s.TokenText() == "tableau_de" {
			tab_type += s.TokenText() + " "
		} else {
			return tab_type + s.TokenText()
		}
	}
}

// func scanDeepExpr(s scanner.Scanner, startRune, endRune rune) string {

// }
