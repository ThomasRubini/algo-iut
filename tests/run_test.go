package tests

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler"
	"algo-iut-1/internal/utils/nopwritecloser"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testOneSyntax(filepath string) (err error) {
	codeBytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
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
			err := testOneSyntax("syntax/" + entry.Name())
			if err == nil {
				t.Fatal("expected error")
			}
		})
	}
}