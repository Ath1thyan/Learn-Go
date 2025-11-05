package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Map structure: line -> (filename -> count in that file)
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			countLines(f, filename, counts)
			f.Close()
		}
	}

	// Print duplicates with file names
	for line, fileCounts := range counts {
		total := 0
		for _, count := range fileCounts {
			total += count
		}
		if total > 1 {
			fmt.Printf("%d\t%s", total, line)
			for filename, count := range fileCounts {
				fmt.Printf("\t[%s:%d]", filename, count)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][filename]++
	}
}
