package main

import "fmt"

func main() {
	var name string

	name = "Lukman Sanjaya"
	fmt.Println(name)

	name = "Lukman S"
	fmt.Println(name)

	var address = "Wonogiri"
	var age = 18

	fmt.Println(address)
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	country = "ID"
	fmt.Println(country)

	// Multiple variable
	var (
		firstName = "Lukman"
		lastName  = "Sanjaya"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
