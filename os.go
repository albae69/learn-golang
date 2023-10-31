package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Args", args)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println("Hostname: ", hostname)

}
