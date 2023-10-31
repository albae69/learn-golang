package main

import "fmt"

func main() {
	var name = ""

	if name == "Bae" {
		fmt.Println("Hello Bae")
	} else if name == "Haqi" {
		fmt.Println("Hello Haqi")
	} else {
		fmt.Println("Who are you?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Wow, your name is longer than 5")
	} else {
		fmt.Println("Hello")
	}
}
