package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	langs := make(map[string]string)

	langs["js"] = "Javascript"
	langs["rb"] = "Ruby"
	langs["py"] = "Python"

	fmt.Println(langs)
	fmt.Println(langs["js"])

	delete(langs, "rb")
	fmt.Println(langs)

	// Loops

	for key, value := range langs {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
