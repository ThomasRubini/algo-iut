package entrypoint

import (
	"algo-iut-1/internal/langoutput"
	"algo-iut-1/internal/scan"
	"algo-iut-1/internal/transpiler"
	"flag"
	"os"
)

func readFileToString(path string) string {

	src, err := os.ReadFile(path)
	if err != nil {
		panic("Error reading file: " + err.Error())
	}

	return string(src)
}

func handleOutput(outputArg string) langoutput.T {
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

	src := readFileToString(*inputArg)
	output := handleOutput(*outputArg)
	defer output.Close()

	s := scan.New(src)
	transpiler.Do(s, output, src)
}
