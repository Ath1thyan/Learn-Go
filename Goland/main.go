package main

import (
	"fmt"
	"net/http"
)

var taskOne = "First Task"
var taskTwo = "Second Task"
var taskThree = "Third Task"

//itemCount := 10

var taskItems = []string{taskOne, taskTwo, taskThree}

func main() {

	//var taskOne = "First Task"
	//var taskTwo = "Second Task"
	//var taskThree = "Third Task"
	//itemCount := 10

	//taskItems := []string{taskOne, taskTwo, taskThree}

	//fmt.Println("###### Hello World! ######")
	//fmt.Println(taskOne)
	//fmt.Println(taskTwo)
	//fmt.Println(taskThree)

	//fmt.Println()
	//fmt.Println("Group Tasks")
	//fmt.Println(taskOne)
	//fmt.Println(taskTwo)

	//fmt.Println()
	//fmt.Println("New Tasks")
	//fmt.Println(taskThree)

	//fmt.Println("Tasks:", taskItems)
	//
	//for _, task := range taskItems {
	//	fmt.Println(task)
	//}

	//printTasks(taskItems)
	//fmt.Println()
	//taskItems = addTask(taskItems, "Hello New Task")
	//printTasks(taskItems)

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	var greeting = "Hello User, welcome to GoLang World!"
	fmt.Fprintln(w, greeting)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	greeting := "Hello World!"
	fmt.Fprintln(w, greeting)
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(w, task)
	}
}

func printTasks(taskItems []string) {
	fmt.Println("###### Todo list ######")
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
	//fmt.Println(itemsLimit)
}

func addTask(taskItems []string, newTask string) []string {
	updatedTaskItems := append(taskItems, newTask)
	return updatedTaskItems
}
