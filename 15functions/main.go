package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greet()
	greet2()

	result, message := adder(10, 5)
	fmt.Println("Adder redult: ", result, message)
	total := proAdd(10, 5, 3, 2, 1)
	fmt.Println("ProAdd redult: ", total)
}

func greet() {
	fmt.Println("Hello")
}
func greet2() {
	fmt.Println("World")
}

func adder(a int, b int) (int, string) {
	return a + b, "Sum"
}

func proAdd(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
