package tests

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler"
	"algo-iut/internal/utils/nopwritecloser"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testOneSyntax(filepath string) (err error) {
	codeBytes, err := os.ReadFile(filepath)
	if err != nil {
		// this is an error in the test itself, not in the transpiler
		panic(fmt.Errorf("read file error: %v", err))
	}

	src := string(codeBytes)
	lang_output := bytes.Buffer{}

	transpile_err := transpiler.Do(scan.New(src), langoutput.NewWriteCloser(nopwritecloser.New(&lang_output)), src)
	if transpile_err != nil {
		return fmt.Errorf("transpile error: %v", transpile_err)
	}

	return nil
}

func TestSyntax(t *testing.T) {
	assert.Equal(t, 1, 1)

	entries, err := os.ReadDir("syntax/")
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			err := testOneSyntax("syntax/" + entry.Name())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestSyntaxFail(t *testing.T) {
	assert.Equal(t, 1, 1)

	entries, err := os.ReadDir("syntax_fail/")
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			err := testOneSyntax("syntax_fail/" + entry.Name())
			if err == nil {
				t.Fatal("expected error")
			}
		})
	}
}

// same as testing the syntax, but these are actual real-world examples, rather than a specific syntax being tested.
// should not be useful, but might as well test :)
func TestExamples(t *testing.T) {
	assert.Equal(t, 1, 1)

	entries, err := os.ReadDir("examples/")
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			err := testOneSyntax("examples/" + entry.Name())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
