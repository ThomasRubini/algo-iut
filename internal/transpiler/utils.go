package transpiler

import (
	"fmt"
	"text/scanner"
)

func PanicInvalidToken(s *scanner.Scanner) {
	panic(fmt.Sprintf("Invalid token: '%s'. Position: %s", s.TokenText(), s.Pos()))
}

func scanType(s *scanner.Scanner) string {
	tab_type := ""

	for {
		if s.TokenText() == "tableau_de" {
			tab_type += s.TokenText() + " "
			s.Scan()
		} else {
			defer s.Scan()
			return tab_type + s.TokenText()
		}
	}
}
