package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpLukman NoKTP = "111222333444"
	var marriedSatatus Married = false

	fmt.Println(noKtpLukman)
	fmt.Println(marriedSatatus)
}
