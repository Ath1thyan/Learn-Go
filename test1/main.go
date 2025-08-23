package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	result := calc(1, 2)
	fmt.Println(result)
}

func calc(x, y int) int {
	res := x + y
	return res
}
