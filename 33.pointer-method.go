package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	lukman := Man{"Lukman"}
	lukman.Married()
	fmt.Println(lukman.Name) // tidak berubah

	lukman.MarriedPointer()
	fmt.Println(lukman.Name) // berubah
}
