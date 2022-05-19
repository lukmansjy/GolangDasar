package database

import "fmt"

var connection string

// func init otomatis dijalankan pertamakali
func init() {
	fmt.Println("function init jalan")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
