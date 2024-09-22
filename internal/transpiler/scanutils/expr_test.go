package scanutils

import (
	"strings"
	"testing"
	"text/scanner"

	"github.com/stretchr/testify/assert"
)

func newScanner(str string) *scanner.Scanner {
	s := scanner.Scanner{}
	s.Init(strings.NewReader(str))
	s.Scan()
	return &s
}

func TestConstExpr(t *testing.T) {
	s := newScanner("1")
	assert.Equal(t, []string{"1"}, Expr(s))
}

func TestAddition(t *testing.T) {
	s := newScanner("1+1")
	assert.Equal(t, []string{"1", "+", "1"}, Expr(s))
}

func TestFuncCall(t *testing.T) {
	s := newScanner("foo(1)")
	assert.Equal(t, []string{"foo", "(", "1", ")"}, Expr(s))

	s = newScanner("foo(1, 5)")
	assert.Equal(t, []string{"foo", "(", "1", ",", "5", ")"}, Expr(s))
}

func TestComplex1(t *testing.T) {
	s := newScanner("foo(1+5, foo2(1, foo3(4), 5))")
	assert.Equal(t, []string{"foo", "(", "1", "+", "5", ",", "foo2", "(", "1", ",", "foo3", "(", "4", ")", "," , "5", ")", ")"}, Expr(s))
}
