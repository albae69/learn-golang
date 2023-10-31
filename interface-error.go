package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := pembagi(100, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Hasil ", hasil)

}
