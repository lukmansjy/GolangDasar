package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Lukman"
	names[1] = "Sanjaya"
	names[2] = "Mantap"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [4]int{
		10,
		80,
		90,
		30,
	}

	fmt.Println(values)
	fmt.Println(values[0])

	values[0] = 78
	fmt.Println(values[0])

	fmt.Println(len(values))
	fmt.Println(len(names))

	var arrLagi [10]string
	fmt.Println(len(arrLagi))
}
