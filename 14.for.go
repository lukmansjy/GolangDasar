package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Ke", counter)
	}

	names := []string{"Lukman", "Sanjaya", "Oke"}

	for i, name := range names {
		fmt.Println("index:", i, "name:", name)
	}

	for _, name := range names {
		fmt.Println("name:", name)
	}
}
