package main

import (
	"flag"
	"fmt"
)

// go run 38.package-flag.go -host=localhost -user=root -password=passwd
func main() {
	host := flag.String("host", "localhost", "Put your database host")
	user := flag.String("user", "root", "Put your database user")
	password := flag.String("password", "pass", "Put your database password")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
}
