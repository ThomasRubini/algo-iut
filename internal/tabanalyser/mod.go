package tabanalyser

import (
	"strings"
)

func countAndRemovePrefix(src string, prefix string) (string, int) {
	n := 0
	for {
		after, found := strings.CutPrefix(src, prefix)
		if !found {
			return src, n
		}

		src = after
		n++
	}
}

func doLine(line string) string {
	prefix := ""

	line, n := countAndRemovePrefix(line, " ")
	prefix += strings.Repeat(" ", n)

	_, n = countAndRemovePrefix(line, "\t")
	prefix += strings.Repeat("\t", n)

	return prefix
}

func Do(src string) []string {
	prefixes := make([]string, 0)
	for _, line := range strings.Split(src, "\n") {
		prefix := doLine(line)
		prefixes = append(prefixes, prefix)
	}
	return prefixes
}
