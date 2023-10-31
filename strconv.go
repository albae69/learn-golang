package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(boolean)

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	fmt.Println(number)

}
