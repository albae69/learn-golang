package main

import "fmt"

func main() {
	// map is key-value based data type
	var person = map[string]string{
		"name":    "Bae",
		"address": "Yogyakarta",
	}
	// add another key
	person["isMarried"] = "Not Yet"

	fmt.Println(person)
	fmt.Println(person["name"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Bae"
	book["wrong"] = "ups"

	fmt.Println(len(book))

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))

}
