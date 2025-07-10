package main

import "fmt"

func main() {
	fmt.Println("Structs")

	athi := User{"Athi", "athi@mail.com", true, 24}
	fmt.Println(athi)
	fmt.Printf("athi details are %+v\n", athi)
	fmt.Printf("Name: %v and email: %v\n", athi.Name, athi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
