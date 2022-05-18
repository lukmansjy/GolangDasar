package main

import (
	"fmt"
)

// type assertion merumakan kemampuan untuk merubah tipe data nebhadu tipe data yg diinginkan

func random() interface{} {
	return "Lukman"
}

func main() {
	// result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // panic error
	// fmt.Println(resultInt)

	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown Tye")
	}
}
