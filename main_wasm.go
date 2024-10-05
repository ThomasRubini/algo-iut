//go:build wasm

package main

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler"
	"algo-iut-1/internal/utils/nopwritecloser"
	"bytes"
	"fmt"
	"io"
	"os"
	"syscall/js"
)

func Wait() {
	c := make(chan struct{})
	<-c
}

// https://stackoverflow.com/a/10476304
func captureStdout(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// Call function
	f()
	w.Close()

	// back to normal state
	os.Stdout = old // restoring the real stdout
	return <-outC
}

func transpile(input string) (output string, logs string, success bool) {
	defer func() {
		if r := recover(); r != nil {
			output = ""
			fmt.Printf("Panicked: %v\n", r)
		}
	}()

	scanner := scan.New(input)
	buf := bytes.Buffer{}

	stdout := captureStdout(func() {

		success = transpiler.Do(
			scanner,
			langoutput.NewWriteCloser(nopwritecloser.New(&buf)),
			input,
		)
	})

	return buf.String(), stdout, success
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
