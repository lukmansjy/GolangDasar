package main

import "fmt"

/*
# parameter di golang secara default pass by value (copy data dan dikirm ke function tsb), ini akan membuat
	variabel asli manjadi aman karena tidak bisa di ubah didalam function
# dan supaya dapat mengubah vaiabel asli kita dapat menggunakan pointer pada parameter. dan untuk menggunakan
	pointer di parameter kita dapat menggunakan operator * di parameter
*/

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

// function pointer
func ChangeAddressToIndonesia2(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	address := Address{"Jogja", "DIY", "ID"}
	ChangeAddressToIndonesia(address)
	fmt.Println(address) // tidak berubah

	ChangeAddressToIndonesia2(&address) // passing sebagai pointer
	fmt.Println(address)                // berubah
}
