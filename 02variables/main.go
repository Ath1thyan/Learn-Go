package main

import "fmt"

func main() {
	var username string = "athi"
	fmt.Println(username)
	fmt.Printf("Variable if of type: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable if of type: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable if of type: %T \n", smallVal)

	var smallFloat float64 = 255.45554511521425874885
	fmt.Println(smallFloat)
	fmt.Printf("Variable if of type: %T \n", smallFloat)

	// Default Values and Aliases

	var newVar int
	fmt.Println(newVar)
	fmt.Printf("Variable if of type: %T \n", newVar)

	var newStr string
	fmt.Println(newStr)
	fmt.Printf("Variable if of type: %T \n", newStr)

	// implicit type

	var website = "athi.com"
	fmt.Println(website)
	fmt.Printf("Variable if of type: %T \n", website)

	// No var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)
	fmt.Printf("Variable if of type: %T \n", numberOfUser)

}
