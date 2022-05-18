package main

import "fmt"

/*
secara default di golang saat inisialisai variabel maka data tersebut akan terisi default value,
misalnya kalau integer 0, boolean false, string berisi string kosong.
ini berbeda dengan bahasa pemrograman lain yg biasanya berus null atau nil
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	person := NewMap("Lukman")
	// dengan nil akan mempermudah penggunaan if, daripada harus cek dengan person["name"] == ""
	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}
