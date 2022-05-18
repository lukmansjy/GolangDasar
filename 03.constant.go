package main

import "fmt"

func main() {
	const firstName string = "Lukman"
	const lastName = "Sanjaya"
	const value = 100

	fmt.Println(firstName)

	// Multiple constant
	const (
		fullName string = "Lukman Sanjaya"
		address         = "Wonogiri"
		value2          = 1000
	)

	fmt.Println(fullName)
}
