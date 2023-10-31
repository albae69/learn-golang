package main

import "fmt"

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Message", message)
	}
	fmt.Println("App stopped")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application error!")
	}
	fmt.Println("Application running...")
}

func main() {
	runApp(false)
}
