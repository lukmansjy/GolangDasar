package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var lukman Customer
	lukman.Name = "Lukman Sanjaya"
	lukman.Address = "Wonogiri"
	lukman.Age = 18

	fmt.Println(lukman)

	// struct literals
	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	// struct literals (best practices)
	joko := Customer{
		Name:    "Joko P",
		Address: "Wonogiri",
		Age:     17,
	}
	fmt.Println(joko)

}
