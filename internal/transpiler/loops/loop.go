package loops

import (
	"io"
	"text/scanner"
)

func DoInfiniteLoop(s *scanner.Scanner, output io.WriteCloser) {
	output.Write([]byte("while(true) {"))
}
