package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

// pointer in function
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // 1.by default it pass the data by value
	var address2 *Address = &address1 // 2.added & to pointer to address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	fmt.Println(address1) // 1.{Subang Jawa Barat Indonesia}
	fmt.Println(address2) // 1.{Bandung Jawa Barat Indonesia} 2. &{Bandung Jawa Barat Indonesia}
	fmt.Println(address3)

	*address2 = Address{"Medan", "Sumatera Utara", "Indonesia"}
	fmt.Println(address2)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	// 1.since it pass by value, when address 2 change, the data from address 1 is not changing.

	// pointer in function
	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat) // add & in parameter since its a pointer
	fmt.Println(alamat)
}
