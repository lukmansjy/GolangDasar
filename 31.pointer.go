package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Wonogiri", "Jawa Tengah", "Indonesia"}
	address2 := address1

	address2.City = "Solo"
	fmt.Println(address1) // Tidak berbuah karena pass by value (datanya di copy)
	fmt.Println(address2)

	// pointer
	address3 := Address{"Semarang", "Jawa Tengah", "Indonesia"}
	address4 := &address3 // ponter (refer ke addess3)

	address4.City = "Jogja"
	fmt.Println(address3) // City berbuah karena pass by reference
	fmt.Println(address4)

	// kalau address4 full di ubah ke address lain maiak address 3 tidak ikuk berubah
	// address4 = &Address{"Cikupa", "Tangerang", "Indonesia"}
	// fmt.Println(address3)
	// fmt.Println(address4)

	// kalau ingin mamaksa mengubah semua
	*address4 = Address{"Balaraja", "Tangerang", "Indonesia"}
	fmt.Println(address3)
	fmt.Println(address4)

	/*
		NOTE:
		* operator & untuk membuat pointer dari data baru ataua variabel yang sudah ada
	*/

	// Pointer kosong
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.City = "Jakarta"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
