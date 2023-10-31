package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ahmad Albaihaqi Lubis", "haqi"))
	fmt.Println(strings.Contains("Ahmad Albaihaqi Lubis", "bae"))
	fmt.Println(strings.Split("Ahmad Albaihaqi Lubis", " "))
	fmt.Println(strings.Split("Ahmad Albaihaqi Lubis", " ")[1])
	fmt.Println(strings.ToLower("Ahmad Albaihaqi Lubis"))
	fmt.Println(strings.ToUpper("Ahmad Albaihaqi Lubis"))
	// fmt.Println(strings.Trim(" a      Ahmad Albaihaqi Lubis b ", ""))
	fmt.Println(strings.ReplaceAll("Ahmad Albaihaqi Lubis", "Ahmad", "Muhammad"))

}
