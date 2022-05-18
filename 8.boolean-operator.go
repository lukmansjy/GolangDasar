package main

import "fmt"

func main() {
	var ujian = 90
	var absesnsi = 80

	var luluUjian = ujian >= 80
	var lulusAbsensi = absesnsi >= 80

	var lulus = lulusAbsensi && luluUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absesnsi >= 80)
}
