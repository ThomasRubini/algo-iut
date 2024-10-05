package scan

import (
	"algo-iut/internal/scan/scanexpr"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstExpr(t *testing.T) {
	s := New("1")
	assert.Equal(t, scanexpr.Id("1"), s.Expr())
}

func TestAddition(t *testing.T) {
	s := New("1+1")
	assert.Equal(t, scanexpr.Merge(scanexpr.Id("1"), scanexpr.Op("+"), scanexpr.Id("1")), s.Expr())
}

func TestFuncCall(t *testing.T) {
	s := New("foo(1)")
	assert.Equal(t, scanexpr.Func("foo", scanexpr.Id("1")), s.Expr())

	s = New("foo(1, 5)")
	assert.Equal(t, scanexpr.Func("foo", scanexpr.Id("1"), scanexpr.Id("5")), s.Expr())
}

func TestComplex1(t *testing.T) {
	s := New("foo(1+5, foo2(1, foo3(4), 5))")
	assert.Equal(t,
		scanexpr.Func("foo",
			scanexpr.Merge(
				scanexpr.Id("1"),
				scanexpr.Op("+"),
				scanexpr.Id("5"),
			),
			scanexpr.Func("foo2",
				scanexpr.Id("1"),
				scanexpr.Func("foo3",
					scanexpr.Id("4"),
				),
				scanexpr.Id("5"),
			),
		), s.Expr())
}

func TestArray(t *testing.T) {
	s := New("arr[1]")
	assert.Equal(t, scanexpr.Arr("arr", scanexpr.Id("1")), s.Expr())
}

func TestEqual(t *testing.T) {
	s := New("1 ==1")
	assert.Equal(t, scanexpr.Merge(
		scanexpr.Id("1"),
		scanexpr.Op("=="),
		scanexpr.Id("1"),
	), s.Expr())
}
func TestTextOperator(t *testing.T) {
	s := New("1 ne_vaut_pas 1")
	assert.Equal(t, scanexpr.Merge(
		scanexpr.Id("1"),
		scanexpr.Op("ne_vaut_pas"),
		scanexpr.Id("1"),
	), s.Expr())
}

func TestTextOperator2(t *testing.T) {
	s := New("tab[j] ne_vaut_pas 0")
	assert.Equal(t, scanexpr.Merge(
		scanexpr.Arr("tab", scanexpr.Id("j")),
		scanexpr.Op("ne_vaut_pas"),
		scanexpr.Id("0"),
	), s.Expr())
}
