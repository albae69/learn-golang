package main

import "fmt"

func main() {
	// slice data type is a slice of an array
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[4:7]
	fmt.Println(slice1)      // [Mei Juni Juli]
	fmt.Println(len(slice1)) // 3
	fmt.Println(cap(slice1)) // 8

	// bulan[5] = "Diubah"
	// fmt.Println(slice1)

	var slice2 = bulan[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bae")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"

	fmt.Println(slice2)
	fmt.Println(bulan)

	// create a slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Bae"
	newSlice[1] = "Lubis"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
