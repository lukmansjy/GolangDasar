package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// with param
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// with return value
func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	}
	return "Hello " + name
}

// multi return value
func getFullName() (string, string) {
	return "Lukman", "Sanjaya"
}

// named return valeus (return berdasarkan variable value)
func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Lukman"
	middleName = "Sanjaya"
	lastName = "Ok"

	return // retrun tidak perlu nyebut variable satu2
}

func main() {
	sayHello()
	sayHelloTo("Lukman", "Sanjaya")

	hello := getHello("Lukman")
	fmt.Println(hello)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName2, _ := getFullName()
	fmt.Println(firstName2)

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
