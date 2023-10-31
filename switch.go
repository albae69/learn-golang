package main

import "fmt"

func main() {
	name := "Bae"

	switch name {
	case "Bae":
		fmt.Println("Hello Bae")
		fmt.Println("Run, Run, Run!")
	default:
		fmt.Println("Sape lu?")

	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Halo")
	}
}
