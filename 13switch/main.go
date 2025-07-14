package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and Case")
	rand.Seed(time.Now().Unix())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Val of dice: ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("1 spot")
	case 2:
		fmt.Println("2 spot")
	case 3:
		fmt.Println("3 spot")
		fallthrough
	case 4:
		fmt.Println("4 spot")
		fallthrough
	case 5:
		fmt.Println("5 spot")
	case 6:
		fmt.Println("6 spot")
	default:
		fmt.Println("Nothing!")
	}
}
