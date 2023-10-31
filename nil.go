package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	person := NewMap("")
	person2 := NewMap("Bae")

	if person == nil {
		fmt.Println("person has no data")
	}

	fmt.Println(person)
	fmt.Println(person2)
}
