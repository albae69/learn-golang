package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome ", name)
	}
}

func blockedUser(name string) bool {
	return name == "Bae"
}

func main() {

	registerUser("Bae", blockedUser)
	registerUser("Haqi", func(s string) bool {
		return s == "Lubis"
	})
}
