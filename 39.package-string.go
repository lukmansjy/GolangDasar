package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Lukman Sanjaya", "Lukman"))
	fmt.Println(strings.Split("Lukman Sanjaya", " "))
	fmt.Println(strings.ToLower("Lukman Sanjaya"))
	fmt.Println(strings.ToUpper("Lukman Sanjaya"))
	fmt.Println(strings.Trim("    Lukman Sanjaya    ", " "))
	fmt.Println(strings.ReplaceAll("Lukman Sanjaya Lukman sanjaya lukaman", "Lukman", "Budi"))
}
