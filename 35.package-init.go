package main

import (
	"GolangDasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
