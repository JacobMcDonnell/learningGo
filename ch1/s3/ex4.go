// prints the count of lines that appear more that once
// handles standard input and files
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	countsFiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countsFiles, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsFiles, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, countsFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsFiles map[string]string, fname string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		var i int = 0
		text := input.Text()
		counts[text]++
		if counts[text] > 1 {
			var s string = countsFiles[text]
			for _, file := range strings.Split(s, "\t") {
				if fname == file {
					i++
				}
			}
			if i == 0 {
				countsFiles[text] += "\t" + fname
			}
		} else {
			countsFiles[text] += "\t" + fname
		}
	}
	// NOTE: ignoring ptential errors from input.Err()
}
