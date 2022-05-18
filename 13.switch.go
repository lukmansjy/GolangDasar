package main

import "fmt"

func main() {
	name := "Lukman"

	switch name {
	case "Lukman":
		fmt.Println("Hello lukman")
	case "Joko":
		fmt.Println("Hello joko")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch kondisi di dalam case
	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Nama terlalu panjang")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
