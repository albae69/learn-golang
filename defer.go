package main

import "fmt"

func logging() {
	fmt.Println("function called")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")

	result := 10 / value
	fmt.Println("Result ", result)

}

func main() {
	runApplication(0)
}
