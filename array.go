package main

import "fmt"

func main() {
	// array is a list of same type data, start with index 0
	var names [3]string // define an array with max length of 3

	names[0] = "Ahmad"
	names[1] = "Albaihaqi"
	names[2] = "Lubis"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
}
