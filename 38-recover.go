package main

import (
	"errors"
	"fmt"
)

// recover() harus ditaro di func yang nantinya dipanggil pake defer

func validateName(name string) (bool, error) {
	if name == "" {
		return false, errors.New("Nama gak boleh kosong yaa!")
	}

	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Ada error nih:", r)
	} else {
		fmt.Println("Yippie, gak ada error")
	}
}

func main() {
	defer catch()

	name := "Rasyidev"
	if valid, err := validateName(name); valid {
		fmt.Println("Halo", name)
	} else {
		panic(err.Error())
	}
}
