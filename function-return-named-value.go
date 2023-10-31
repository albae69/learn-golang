package main

import "fmt"

func getFullName2() (firstName string, lastName string) {
	firstName = "Bae"
	lastName = "Lubis"
	return
}

func main() {
	firstName, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
