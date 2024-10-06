package langoutput

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	o := NewString()
	o.Write("hello")
	assert.Equal(t, "hello", o.String())
}
