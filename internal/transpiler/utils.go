package transpiler

import (
	"text/scanner"
)

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
