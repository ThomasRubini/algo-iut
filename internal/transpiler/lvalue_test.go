package transpiler

import (
	"algo-iut/internal/scan"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionArgsEmpty(t *testing.T) {
	args := getFunctionArgs(scan.New(")"))
	assert.Equal(t, []string{}, args)
}

func TestFunctionArgsNormal(t *testing.T) {
	args := getFunctionArgs(scan.New("5, 6, 7)"))
	assert.Equal(t, []string{"5", "6", "7"}, args)
}

func TestFunctionArgsExpr(t *testing.T) {
	args := getFunctionArgs(scan.New("5+5, (6+1), 7)"))
	assert.Equal(t, []string{"5 + 5", "( 6 + 1 )", "7"}, args)
}
