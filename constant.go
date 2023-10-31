package main

import "fmt"

func main() {
	// constant is create a constant variable that the value never change
	// it's okay if the variable not called
	const (
		firstName = "Ahmad"
		midName   = "Albaihaqi"
		lastName  = "Lubis"
	)

	fmt.Println("firstName", firstName)
	// fmt.Println("midName", midName)
	fmt.Println("lastName", lastName)
}
