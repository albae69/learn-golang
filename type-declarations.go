package main

import "fmt"

func main() {
	// type is an alias for data type
	type NIK string

	var personNik NIK = "123456789011"
	fmt.Println(personNik)
}
