package transpiler

import "algo-iut-1/internal/scan"

func scanType(s scan.Scanner) string {
	tab_type := ""

	for {
		if s.Peek() == "tableau_de" {
			tab_type += s.Peek() + " "
			s.Advance()
		} else {
			defer s.Advance()
			return tab_type + s.Peek()
		}
	}
}
