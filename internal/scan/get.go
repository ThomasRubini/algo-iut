package scan

import (
	"fmt"
	"slices"
	"strconv"
)

func (s *impl) Text() string {
	str := s.Peek()
	s.Advance()
	return str
}

func (s *impl) Number() int {
	str := s.Text()
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to int: %v", str, err))
	}
	return num
}

var keywords = []string{
	// loops
	"tant_que", "pour", "boucle", "fboucle",
	"faire", "ffaire", "sortie",
	// function stuff
	"fonction", "procedure", "renvoie",
	// others
	"debut","fin",
}

func (s *impl) LValue() string {
	tok := s.Text()
	if slices.Contains(keywords, tok) {
		s.InvalidToken("Expected lvalue, found reserved keyword")
	}

	if s.Peek() == "[" {
		s.Text() // consume '['
		inside := s.LValue()
		s.Must("]")
		return fmt.Sprintf("%v[%v]", tok, inside)
	} else {
		return tok
	}
}
