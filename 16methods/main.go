package main

import "fmt"

func main() {
	fmt.Println("Methods")

	athi := User{"Athi", "athi@mail.com", true, 24}
	fmt.Println(athi)
	fmt.Printf("athi details are %+v\n", athi)
	fmt.Printf("Name: %v and email: %v\n", athi.Name, athi.Email)
	athi.GetStatus()
	athi.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@athi.in"
	fmt.Println("user email: ", u.Email)
}
