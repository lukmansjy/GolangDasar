package main

import (
	"fmt"
)

func main() {
	// map[typeKey]typeValue
	person := map[string]string{
		"name":    "Lukman",
		"address": "Wonogiri",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	person["name"] = "Lukman S"
	fmt.Println(person["name"])

	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Belajar Go"
	book["author"] = "Lukman"
	book["publish"] = "2022-03-21"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "publish")
	fmt.Println(book)
	fmt.Println(len(book))
}
