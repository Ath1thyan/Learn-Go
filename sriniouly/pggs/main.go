package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	// fmt.Println("Myfavorite number is", rand.Intn(10))
	fmt.Println(split(5))
}

func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}
