package main

import "fmt"

func main() {
	fmt.Println("Array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Kiwi"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List is: ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushrooom"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list is: ", len(vegList))
}
