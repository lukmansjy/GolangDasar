package main

import "fmt"

func logging() {
	fmt.Println("selesai memangil function")
}

func runAplication(value int) {
	// menambahkah defer, fungsi akan dijalankan terakhir & akan tetap dijalankan walau error
	defer logging()
	fmt.Println("Run application")
	fmt.Println(10 / value) // error panic integer divide by zero
}

func main() {
	runAplication(0)
}
