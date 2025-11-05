package main

import (
	"bufio"
	"os"
	"fmt"
)

// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]

// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// }

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines2(os.Stdin, counts)
	} else {
		for _, args := range files {
			file, err := os.Open(args)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines2(file, counts)
			file.Close()
		}
	}
	done(counts)
}

func done(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines2(f *os.File, counts map[string]int) {
	inp := bufio.NewScanner(f)
	for inp.Scan() {
		counts[inp.Text()]++
	}
}