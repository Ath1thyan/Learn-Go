package main

import (
	"fmt"
	// "time"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// fmt.Println("Saturday")
	// fmt.Println(time.Now().Weekday())
	// fmt.Println(time.Saturday - time.Now().Weekday())
	// fmt.Println(time.Now())

	// today := time.Now().Weekday()
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today")
	// case today + 1:
	// 	fmt.Println("Tomorrow")
	// case today + 2:
	// 	fmt.Println("In two days")
	// case today + 3:
	// 	fmt.Println("In three days")
	// case today + 4:
	// 	fmt.Println("In four days")
	// default:
	// 	fmt.Println("Too far away")
	// }

	// t := time.Now()

	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("Good Morning")
	// case t.Hour() < 18:
	// 	fmt.Println("Good Afternoon")
	// default:
	// 	fmt.Println("Good Evening")
	// }

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

}
