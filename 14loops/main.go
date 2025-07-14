package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	fmt.Println("-----------------------------")

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("-----------------------------")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("-----------------------------")

	for index, day := range days {
		fmt.Printf("Index -> %v, value -> %v\n", index, day)
	}

	fmt.Println("-----------------------------")

	for _, day := range days {
		fmt.Printf("Index -> value -> %v\n", day)
	}

	fmt.Println("-----------------------------")

	rougeVal := 1
	for rougeVal < 10 {
		if rougeVal == 2 {
			goto ath
		}
		if rougeVal == 4 {
			rougeVal++
			continue
		}
		if rougeVal == 6 {
			break
		}
		fmt.Println("Rouge Value: ", rougeVal)
		rougeVal++
	}

ath:
	fmt.Println("Athi")

	fmt.Println("-----------------------------")

}
