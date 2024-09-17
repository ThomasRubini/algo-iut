package entrypoint

import (
	"algo-iut-1/internal/transpiler"
	"os"
	"strings"
	"text/scanner"
)

func readFileToString(path string) string {

	src, err := os.ReadFile(path)
	if err != nil {
		panic("Error reading file: " + err.Error())
	}

	return string(src)
}

func Main() {
	src := readFileToString("input.txt")

	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	transpiler.DoRoot(&s, os.Stdout)

}
