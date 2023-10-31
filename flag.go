package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "put  your database host name")
	var user *string = flag.String("user", "root", "put  your database user")
	var password *string = flag.String("password", "root", "put  your database password")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)

}
