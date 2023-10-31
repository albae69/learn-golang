package main

import "fmt"

func main() {
	// variable is like a container that has a space to save some data
	var name string

	name = "Bae"
	fmt.Println(name)

	name = "Ahmad"
	fmt.Println(name)

	var friend_name = "soso"
	fmt.Println(friend_name)

	var age = 23
	fmt.Println(age)

	first_name := "ahmad"
	fmt.Println(first_name)

	var (
		isMarried = false
		isAlive   = false
	)
	fmt.Println("isMarried", isMarried)
	fmt.Println("isAlive", isAlive)
}
