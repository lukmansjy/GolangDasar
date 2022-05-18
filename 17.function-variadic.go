package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 12, 9, 8, 12, 16)
	fmt.Println(total)

	slice := []int{10, 12, 9, 8, 12, 16}
	total2 := sumAll(slice...)
	fmt.Println(total2)
}
