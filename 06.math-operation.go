package main

import "fmt"

func main() {
	var result = 10 + 15
	fmt.Println(result)

	var a = 10
	var b = 20
	var c = a * b
	fmt.Println(c)

	// Augmented Assignments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	// Unary Operator
	var y = 10
	y++
	fmt.Println(y)
}
