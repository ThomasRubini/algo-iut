package scan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstExpr(t *testing.T) {
	s := New("1")
	assert.Equal(t, []string{"1"}, s.Expr())
}

func TestAddition(t *testing.T) {
	s := New("1+1")
	assert.Equal(t, []string{"1", "+", "1"}, s.Expr())
}

func TestFuncCall(t *testing.T) {
	s := New("foo(1)")
	assert.Equal(t, []string{"foo", "(", "1", ")"}, s.Expr())

	s = New("foo(1, 5)")
	assert.Equal(t, []string{"foo", "(", "1", ",", "5", ")"}, s.Expr())
}

func TestComplex1(t *testing.T) {
	s := New("foo(1+5, foo2(1, foo3(4), 5))")
	assert.Equal(t, []string{"foo", "(", "1", "+", "5", ",", "foo2", "(", "1", ",", "foo3", "(", "4", ")", ",", "5", ")", ")"}, s.Expr())
}

func TestArray(t *testing.T) {
	s := New("arr[1]")
	assert.Equal(t, []string{"arr", "[", "1", "]"}, s.Expr())
}

func TestEqual(t *testing.T) {
	s := New("1 ==1")
	assert.Equal(t, []string{"1", "==", "1"}, s.Expr())
}
