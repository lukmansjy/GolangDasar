package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}

	return value * factorialRecursive(value-1)
}

func main() {
	factorial := factorialRecursive(10)
	fmt.Println(factorial)
}
