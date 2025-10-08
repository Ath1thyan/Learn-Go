package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	benchmarkTime := time.Now()
	var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	time.Sleep(1 * time.Second)
	timeElapsed := time.Since(benchmarkTime)
	fmt.Println(timeElapsed)

	fmt.Println("-----------------------------")

	benchmarkTime2 := time.Now()
	e, eep := "", ""
	for _, arg := range os.Args[1:] {
		e += eep + arg
		eep = " "
	}
	fmt.Println(e)
	time.Sleep(1 * time.Second)
	timeElapsed2 := time.Since(benchmarkTime2)
	fmt.Println(timeElapsed2)
	fmt.Println("-----------------------------")

	benchmarkTime3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	time.Sleep(1 * time.Second)
	timeElapsed3 := time.Since(benchmarkTime3)
	fmt.Println(timeElapsed3)
	fmt.Println("-----------------------------")

	benchmarkTime4 := time.Now()
	fmt.Println(os.Args[1:])
	time.Sleep(1 * time.Second)
	timeElapsed4 := time.Since(benchmarkTime4)
	fmt.Println(timeElapsed4)
	fmt.Println("-----------------------------")

	fmt.Println(os.Args[0])
	fmt.Println(os.Args)

	fmt.Println(len(os.Args))
	fmt.Println("-----------------------------")

	benchmarkTime5 := time.Now()
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	time.Sleep(1 * time.Second)
	timeElapsed5 := time.Since(benchmarkTime5)
	fmt.Println(timeElapsed5)
	fmt.Println("-----------------------------")

}
