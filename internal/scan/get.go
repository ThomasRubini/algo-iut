package scan

import (
	"fmt"
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

func (s *impl) LValue() string {
	tok := s.Text()
	if s.Peek() == "[" {
		s.Text() // consume '['
		inside := s.LValue()
		s.Must("]")
		return fmt.Sprintf("%v[%v]", tok, inside)
	} else {
		return tok
	}
}
