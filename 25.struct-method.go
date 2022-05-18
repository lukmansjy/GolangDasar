package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi() {
	fmt.Println("Hello, My Name is", customer.Name)
}

func main() {
	lukman := Customer{Name: "Lukman Sanjaya"}
	lukman.sayHi()
}
