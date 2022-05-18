package main

import "fmt"

func main() {
	name := "Lukman"

	if name == "Lukman" {
		fmt.Println("Hello Lukman")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, boleh kenalan")
	}

	fmt.Println(len(name))

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
