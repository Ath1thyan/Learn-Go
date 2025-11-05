package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// func dup1() {
// 	count := make(map[string]int)
// 	inp := bufio.NewScanner(os.Stdin)

// 	for inp.Scan() {
// 		count[inp.Text()]++
// 	}

// 	for new, n := range count {
// 		if n > 1 {
// 			fmt.Printf("%d \t %s \n", n, new)
// 		}
// 	}
// }