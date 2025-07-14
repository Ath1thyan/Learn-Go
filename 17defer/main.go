package main

import "fmt"

func main() {
	fmt.Println("Defer")
	defer fmt.Println("World0")
	defer fmt.Println("World1")
	myDefer()
	defer fmt.Println("World2")
	fmt.Println("Hello")
	defer fmt.Println("World3")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
