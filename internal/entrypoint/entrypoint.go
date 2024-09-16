package entrypoint

import (
	"algo-iut-1/internal/transpiler"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

func Main() {
	src, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var s scanner.Scanner
	s.Init(strings.NewReader(string(src)))

	transpiler.DoRoot(&s, os.Stdout)

}
