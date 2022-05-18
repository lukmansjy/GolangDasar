package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error message:", message)
	}
	fmt.Println("APlikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("test print")
}
