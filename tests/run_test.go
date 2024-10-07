package tests

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"path/filepath"

	"github.com/stretchr/testify/assert"
)

// recursively read directory
func readDir(dirPath string) []string {
	var filepaths []string

	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			filepaths = append(filepaths, path)
		}
		return nil
	})

	if err != nil {
		panic(fmt.Errorf("walk error: %v", err))
	}

	return filepaths
}

func testOneSyntax(filepath string) (cpp_data string, err error) {
	codeBytes, err := os.ReadFile(filepath)
	if err != nil {
		// this is an error in the test itself, not in the transpiler
		panic(fmt.Errorf("read file error: %v", err))
	}

	src := string(codeBytes)
	lang_output := langoutput.NewString()

	transpile_err := transpiler.Do(scan.New(src), lang_output, src)
	if transpile_err != nil {
		return "", fmt.Errorf("transpile error: %v", transpile_err)
	}

	return lang_output.String(), nil
}

func checkOneCpp(cpp_data string) (err error) {
	cmd := exec.Command("g++", "-x", "c++", "-o", "/dev/null", "-")
	cmd.Stdin = strings.NewReader(cpp_data)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gcc error: %v", err)
	}

	return nil
}

func TestSyntax(t *testing.T) {
	assert.Equal(t, 1, 1)

	for _, entry := range readDir("syntax/") {
		t.Run(entry, func(t *testing.T) {
			_, err := testOneSyntax(entry)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestSyntaxFail(t *testing.T) {
	assert.Equal(t, 1, 1)

	for _, entry := range readDir("syntax_fail/") {
		t.Run(entry, func(t *testing.T) {
			_, err := testOneSyntax(entry)
			if err == nil {
				t.Fatal("expected error")
			}
		})
	}
}

// same as testing the syntax, but these are actual real-world examples, rather than a specific syntax being tested.
// should not be useful, but might as well test :)
// +, we test output C++ syntax validity
func TestExamples(t *testing.T) {
	assert.Equal(t, 1, 1)

	for _, entry := range readDir("examples/") {
		t.Run(entry, func(t *testing.T) {
			var cpp_data string
			t.Run("Transpile", func(t *testing.T) {
				var err error
				cpp_data, err = testOneSyntax(entry)
				if err != nil {
					t.Fatal(err)
				}
			})
			if cpp_data != "" {
				t.Run("CheckCpp", func(t *testing.T) {
					err := checkOneCpp(cpp_data)
					if err != nil {
						t.Fatal(err)
					}
				})
			}
		})
	}
}
