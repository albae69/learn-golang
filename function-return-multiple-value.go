package main

import "fmt"

func getFullName() (string, string) {
	return "Ahmad", "Albaihaqi"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println("First Name : ", firstName)
	fmt.Println("Last Name : ", lastName)
}
