package main

import "fmt"

func main() {
	fmt.Println("If Else")
	loginCount := 10

	var result string
	if loginCount < 0 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "10 login count"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("not less than 10")
	}

}
