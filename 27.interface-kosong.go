package main

import "fmt"

func Ups() interface{} {
	// bisa return type data apapun
	// return 1
	// return true
	return "Ups"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
