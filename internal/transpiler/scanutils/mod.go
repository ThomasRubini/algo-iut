package scanutils

import (
	"fmt"
	"text/scanner"
)

func ScanAndReturn(s *scanner.Scanner) string {
	s.Scan()
	return s.TokenText()
}

func Must(s *scanner.Scanner, str string) {
	s.Scan()
	if s.TokenText() != str {
		panic(fmt.Sprintf("expected %s, got %s", str, s.TokenText()))
	}
}
