package main

import (
	"fmt"
	"learn/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
