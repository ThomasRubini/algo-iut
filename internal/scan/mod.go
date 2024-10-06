package scan

import (
	"algo-iut/internal/scan/scanexpr"
	"fmt"
	"strings"
	"text/scanner"
)

type Scanner interface {
	Peek() string
	Advance()
	Match(string) bool
	Must(string)
	UntilEOL() string
	HasMore() bool

	Text() string
	Number() int
	// value at the left of an assignation. e.g. `tab[5]` in `tab[5] <- 5`
	LValue() string
	Expr() scanexpr.Comp

	InvalidToken(string)
	Pos() scanner.Position

	// best thing to do if we want to rewind the scanner
	Copy() Scanner
}

type impl struct {
	goImpl scanner.Scanner
}

func New(input string) Scanner {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Scan() // first scan
	return &impl{goImpl: s}
}

func (s *impl) HasMore() bool {
	return s.goImpl.Peek() != scanner.EOF
}

func (s *impl) InvalidToken(help string) {
	panic(fmt.Sprintf("Invalid token: '%s'. Position: %s. %s", s.Peek(), s.Pos(), help))
}

func (s *impl) Peek() string {
	return s.goImpl.TokenText()
}

func (s *impl) Advance() {
	s.goImpl.Scan()
}

func (s *impl) Match(str string) bool {
	if s.Peek() == str {
		s.Advance()
		return true
	} else {
		return false
	}
}

func (s *impl) Must(expected string) {
	got := s.Peek()
	if got == expected {
		s.Advance()
	} else {
		panic(fmt.Sprintf("expected '%s', got '%s' (position: %v)", expected, got, s.Pos()))
	}
}

func (s *impl) Pos() scanner.Position {
	return s.goImpl.Pos()
}

func (s *impl) Copy() Scanner {
	clone := *s
	return &clone
}
