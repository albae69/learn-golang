package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsen = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusNilaiAbsen = nilaiAbsen > 80

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsen
	fmt.Println(lulus)
}
