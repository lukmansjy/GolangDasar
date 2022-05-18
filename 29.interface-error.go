package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	}
	return nilai / pembagi, nil
}

func main() {
	result, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", err.Error())
	}

}
