package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	bae := Man{"Bae"}
	bae.Married()

	fmt.Println(bae)
}
