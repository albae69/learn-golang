package main

import "fmt"

type Person struct {
	name      string
	age       int
	isMarried bool
}

// struct method
func (person Person) greeting(name string) {
	fmt.Println("Hello, ", name, " my name is ", person.name)
}

func main() {
	var person Person
	person.name = "Bae"
	person.age = 23
	person.isMarried = false

	fmt.Println(person)
	// calling struct method
	person.greeting("Eko")

	// struct literal
	Tole :=
		Person{
			name:      "Tole",
			age:       21,
			isMarried: true,
		}

	fmt.Println(Tole)
}
