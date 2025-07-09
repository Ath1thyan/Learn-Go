package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var ptr2 *int
	fmt.Println("Val of ptr2 is ", ptr2)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Val of ptr is ", ptr)
	fmt.Println("Val of ptr is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value is ", myNumber)
}
