package main

import "fmt"

func main() {
	angka := -1
	var result string

	if angka < 0 {
		result = "Negatif"
	} else if angka == 0 {
		result = "Nol"
	} else {
		result = "Positif"
	}

	fmt.Println("Angka \t:", angka)
	fmt.Println("Result \t:", result)
}
