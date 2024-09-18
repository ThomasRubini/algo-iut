package scanutils

import (
	"fmt"
	"strconv"
	"text/scanner"
)

func Text(s *scanner.Scanner) string {
	tok := s.TokenText()
	s.Scan()
	return tok
}

func Type(s *scanner.Scanner) string {
	tok := Text(s)
	if tok == "tableau_de" {
		return tok + " " + Type(s)
	} else {
		return tok
	}
}

func UntilEOL(s *scanner.Scanner) string {
	str := ""
	for {
		if s.TokenText() == ";" {
			s.Scan()
			return str[1:]
		} else {
			str += " " + s.TokenText()
			s.Scan()
		}
	}
}

func Number(s *scanner.Scanner) int {
	str := Text(s)
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to int: %v", str, err))
	}
	return num
}

func Must(s *scanner.Scanner, expected string) {
	got := Text(s)
	if got != expected {
		panic(fmt.Sprintf("expected '%s', got '%s' (position: %v)", expected, got, s.Pos()))
	}
}

func Match(s *scanner.Scanner, expected string) bool {
	if s.TokenText() == expected {
		s.Scan()
		return true
	} else {
		return false
	}
}
