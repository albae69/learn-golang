package main

import "fmt"

func main() {
	// for counter := 0; counter <= 10; counter++ {
	// 	fmt.Println("looping ke-", counter)
	// }

	slice := []string{"Ahmad", "Albaihaqi", "Lubis"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	for _, value := range slice {
		fmt.Println(value)

	}
}
