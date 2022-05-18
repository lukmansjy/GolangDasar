package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // menghentikan perulangan
		}
		fmt.Println("Perulangan ke", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip
		}
		fmt.Println(i)
	}
}
