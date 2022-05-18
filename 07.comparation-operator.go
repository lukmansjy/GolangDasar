package main

import "fmt"

func main() {
	var name1 = "Lukman"
	var name2 = "Lukman"

	var result bool = name1 == name2
	fmt.Println(result)

	fmt.Println(name1 == "Budi")
	fmt.Println(10 < 9)
	fmt.Println(10 > 9)
	fmt.Println(10 != 9)
}
