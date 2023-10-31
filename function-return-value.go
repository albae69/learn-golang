package main

import "fmt"

func getName(name string) string {
	if name == "" {
		return "Hello unknown!"
	} else {
		return "Hello " + name
	}

}

func main() {
	result := getName("")

	fmt.Println(result)
}
