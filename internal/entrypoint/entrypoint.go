package entrypoint

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler"
	"flag"
	"fmt"
	"io"
	"os"
)

func readInput(path string) string {

	if path == "-" {
		// Read from stdin
		src, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic("Error reading from stdin: " + err.Error())
		}
		return string(src)
	} else {
		src, err := os.ReadFile(path)
		if err != nil {
			panic("Error reading file: " + err.Error())
		}
		return string(src)
	}
}

func setupOutput(outputArg string) langoutput.T {
	if outputArg == "-" {
		return langoutput.NewWriteCloser(os.Stdout)
	} else {
		output, err := os.Create(outputArg)
		if err != nil {
			panic("Error creating output file: " + err.Error())
		}
		return langoutput.NewWriteCloser(output)
	}
}

func Main() {
	outputArg := flag.String("o", "output.txt", "output file. Use '-' for stdout")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <algo file> [options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults() // Print flag options
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	src := readInput(flag.Arg(0))
	output := setupOutput(*outputArg)
	defer output.Close()

	s := scan.New(src)

	transpile_err := transpiler.Do(s, output, src)
	if transpile_err != nil {
		transpile_err.Show(os.Stderr)
		os.Exit(1)
	}
}
