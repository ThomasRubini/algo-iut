package scanutils

import (
	"fmt"
	"strconv"
	"text/scanner"
)

func ScanAndReturn(s *scanner.Scanner) string {
	s.Scan()
	return s.TokenText()
}

func Number(s *scanner.Scanner) int {
	str := ScanAndReturn(s)
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to int: %v", str, err))
	}
	return num
}

func Must(s *scanner.Scanner, str string) {
	s.Scan()
	if s.TokenText() != str {
		panic(fmt.Sprintf("expected %s, got %s", str, s.TokenText()))
	}
}
