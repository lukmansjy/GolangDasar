package main

import "fmt"

type Filter func(string) string

func sayHalloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "embek" {
		return "...."
	}
	return name
}

func main() {
	sayHalloWithFilter("Lukman", spamFilter)
	sayHalloWithFilter("embek", spamFilter)
}
