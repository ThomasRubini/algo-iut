package utils

import (
	"fmt"
	"text/scanner"
)

func PanicInvalidToken(s *scanner.Scanner, help string) {
	panic(fmt.Sprintf("Invalid token: '%s'. Position: %s. %s", s.TokenText(), s.Pos(), help))
}
