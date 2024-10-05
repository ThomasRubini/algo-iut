package entrypoint

import (
	"algo-iut/internal/langoutput"
	"algo-iut/internal/scan"
	"algo-iut/internal/transpiler"
	"flag"
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

func stringFlag(name, defaultValue, help string) *string {
	var str string
	flag.StringVar(&str, name, defaultValue, help)
	flag.StringVar(&str, string(name[0]), defaultValue, help)
	return &str
}

func Main() {
	inputArg := stringFlag("input", "input.txt", "input file")
	outputArg := stringFlag("output", "output.txt", "output file. Use '-' for stdout")
	flag.Parse()

	src := readInput(*inputArg)
	output := setupOutput(*outputArg)
	defer output.Close()

	s := scan.New(src)

	transpile_err := transpiler.Do(s, output, src)
	if transpile_err != nil {
		transpile_err.Show(os.Stderr)
		os.Exit(1)
	}
}
