//go:build wasm

package main

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler"
	"algo-iut-1/internal/utils/nopwritecloser"
	"bytes"
	"fmt"
	"syscall/js"
)

func Wait() {
	c := make(chan struct{})
	<-c
}

func transpile(input string) (output string, logs string, success bool) {
	defer func() {
		if r := recover(); r != nil {
			output = ""
			fmt.Printf("Panicked: %v\n", r)
		}
	}()

	scanner := scan.New(input)
	lang := bytes.Buffer{}

	transpile_err := transpiler.Do(
		scanner,
		langoutput.NewWriteCloser(nopwritecloser.New(&lang)),
		input,
	)

	logs_buf := bytes.Buffer{}
	if transpile_err != nil {
		transpile_err.Show(&logs_buf)
	}
	success = transpile_err == nil

	return lang.String(), logs_buf.String(), success
}

// see transpile() for signature
func transpileJs(this js.Value, vals []js.Value) any {
	if len(vals) != 1 {
		return "expected 1 argument"
	}
	input := vals[0].String()

	output, logs, success := transpile(input)
	return []interface{}{output, logs, success}
}

func main() {
	fmt.Println("Go program started !")
	js.Global().Set("transpile", js.FuncOf(transpileJs))
	Wait()
}
