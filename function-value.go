package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func main() {
	sayHelloTo := greet
	fmt.Println(sayHelloTo("Bae"))
}
