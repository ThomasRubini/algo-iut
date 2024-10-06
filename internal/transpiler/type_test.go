package transpiler

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeBasic(t *testing.T) {
	o := langoutput.NewString()
	doTypeNoSize(scan.New("entier"), o)
	assert.Equal(t, "int", o.String())
}
func TestTypeTableau(t *testing.T) {
	o := langoutput.NewString()
	doTypeNoSize(scan.New("tableau_de entier"), o)
	assert.Equal(t, "std::vector<int>", o.String())
}
