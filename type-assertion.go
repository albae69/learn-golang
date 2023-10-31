package main

import "fmt"

func random() interface{} {
	return 10
}

func main() {
	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown type!")
	}

	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt int = result.(int)
	// fmt.Println(resultInt)
}
