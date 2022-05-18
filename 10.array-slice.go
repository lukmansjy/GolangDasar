package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "DiubahLagi"
	fmt.Println(months)

	var slice2 = months[10:]
	// var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Lukman")
	fmt.Println(slice3)

	slice3[1] = "BukanDesember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Make new slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Lukman"
	newSlice[1] = "Sanjaya"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perbedaan membuat array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
