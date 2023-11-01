package main

import (
	"fmt"
	"os"
)

// coba jalankan dengan  go run 37.package-os.go lukman sanjaya
func main() {
	args := os.Args
	fmt.Println("Argument: ")
	//fmt.Println(args)
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Hostname", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

}
