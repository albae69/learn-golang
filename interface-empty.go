package main

import "fmt"

type Any interface{}

func ups(i int) Any {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	fmt.Println(ups(3))
}
