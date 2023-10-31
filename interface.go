package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Orang struct {
	Name string
}

func (orang Orang) GetName() string {
	return "Hello, " + orang.Name
}

func main() {
	var bae Orang
	bae.Name = "Bae"

	SayHello(bae)
}
